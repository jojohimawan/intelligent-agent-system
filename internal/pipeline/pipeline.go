package pipeline

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"ev-gitlab.mataelang.net/ev-connect/create-ias/can.git"
	pb "github.com/jojohimawan/intelligent-agent-system/api"
	internalcan "github.com/jojohimawan/intelligent-agent-system/internal/can"
	"github.com/jojohimawan/intelligent-agent-system/internal/kafka"
	"github.com/jojohimawan/intelligent-agent-system/internal/nmea"
	"github.com/jojohimawan/intelligent-agent-system/internal/serial"
	"github.com/jojohimawan/intelligent-agent-system/internal/util"
)

func Run(
	ctx context.Context,
	sr *serial.SerialReader,
	sv *serial.VcanConnection,
	producer *kafka.Producer,
) error {
	rawSentences := make(chan string, 50)
	locations := make(chan *pb.LocationRequest, 50)

	rawFrame := make(chan can.Frame, 50)
	decodedFrame := make(chan *internalcan.DecodedSignal, 50)
	parsedFrame := make(chan *pb.TelematicsBatch, 50)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	decoder := internalcan.NewDecoder()

	var wg sync.WaitGroup

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		select {
		case <-sigCh:
			log.Println("received interrupt signal, shutting down...")
			cancel()
		case <-ctx.Done():
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		ReadCanFrameLoop(ctx, sv, rawFrame)
		close(rawFrame)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DecodeFrameLoop(ctx, decoder, rawFrame, decodedFrame)
		close(decodedFrame)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		ParseFrameLoop(ctx, decodedFrame, parsedFrame)
		close(parsedFrame)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		PublishFrameLoop(ctx, producer, parsedFrame)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		ReadSerialLoop(ctx, sr, rawSentences)
		close(rawSentences)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		ParseLoop(ctx, rawSentences, locations)
		close(locations)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		PublishLoop(ctx, producer, locations)
	}()

	wg.Wait()
	log.Printf("pipeline terminated cleanly.")
	return nil
}

func ReadCanFrameLoop(ctx context.Context, sr *serial.VcanConnection, out chan<- can.Frame) {
	type readResult struct {
		frame can.Frame
		err   error
	}
	resultCh := make(chan readResult)

	go func() {
		defer close(resultCh)

		for sr.Recv.Receive() {
			frame := sr.Recv.Frame()

			log.Println("Received frame...")

			select {
			case resultCh <- readResult{frame, nil}:
			case <-ctx.Done():
				return
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case frame, ok := <-resultCh:
			if !ok {
				return
			}

			out <- frame.frame

		}
	}
}

func DecodeFrameLoop(ctx context.Context, d *internalcan.Decoder, in <-chan can.Frame, out chan<- *internalcan.DecodedSignal) {
	for {
		select {
		case <-ctx.Done():
			return
		case frame, ok := <-in:
			if !ok {
				return
			}

			log.Printf("Pipeline: Processing Frame ID 0x%X", frame.ID)

			signals, err := d.Decode(frame)
			if err != nil {
				log.Printf("Pipeline: Decode Error for ID 0x%X: %v", frame.ID, err)
				continue
			}

			for _, sig := range signals {
				log.Printf("Pipeline: Decoded Signal")
				select {
				case out <- sig:
				case <-ctx.Done():
					return
				}
			}
		}
	}
}

func ParseFrameLoop(ctx context.Context, in <-chan *internalcan.DecodedSignal, out chan<- *pb.TelematicsBatch) {
	const batchSize = 10
	buffer := make([]*internalcan.DecodedSignal, 0, batchSize)
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	flush := func() {
		if len(buffer) == 0 {
			return
		}

		batchMsg, err := internalcan.MarshalSignal("4S4BRDLC3B2413966", buffer)
		if err != nil {
			fmt.Printf("Error marshaling batch: %v\n", err)
		} else {
			select {
			case out <- batchMsg:
			case <-ctx.Done():
				return

			}
		}

		buffer = buffer[:0]
	}

	for {
		select {
		case <-ctx.Done():
			flush()
			return
		case decodedFrame, ok := <-in:
			if !ok {
				flush()
				return
			}

			buffer = append(buffer, decodedFrame)

			if len(buffer) >= batchSize {
				flush()
				ticker.Reset(10 * time.Second)
			}
		case <-ticker.C:
			flush()
		}
	}
}

func PublishFrameLoop(ctx context.Context, producer *kafka.Producer, in <-chan *pb.TelematicsBatch) {
	topic := "vehicle-telematics"

	for {
		select {
		case <-ctx.Done():
			log.Println("Context cancelled, flushing producer...")
			producer.Flush()
			log.Println("Producer flushed. exiting loop...")
			return
		case parsedMsg, ok := <-in:
			if !ok {
				log.Println("Channel closed, flushing producer...")
				producer.Flush()
				log.Println("Producer flushed. exiting loop...")
				return
			}

			if err := producer.PublishTelematics(parsedMsg, &topic); err != nil {
				log.Printf("Kafka publish error: %v", err)
			}
		}
	}
}

func ReadSerialLoop(ctx context.Context, sr *serial.SerialReader, out chan<- string) {
	if sr == nil {
		log.Println("Serial reader is nil, skipping...")
		return
	}

	type readResult struct {
		line string
		err  error
	}
	resultCh := make(chan readResult)

	go func() {
		defer close(resultCh)

		var buffer string

		for {
			line, err := sr.ReadLine()
			if err != nil {
				resultCh <- readResult{"", err}
				continue
			}

			buffer += line

			if strings.Contains(buffer, "\n") {
				parts := strings.Split(buffer, "\n")

				for i := 0; i < len(parts)-1; i++ {
					s := strings.TrimSpace(parts[i])
					if s != "" {
						select {
						case resultCh <- readResult{s, nil}:
						case <-ctx.Done():
							return
						}
					}
				}

				buffer = parts[len(parts)-1]
			}
		}

	}()

	for {
		select {
		case <-ctx.Done():
			return
		case res, ok := <-resultCh:
			if !ok {
				return
			}

			if res.err != nil {
				log.Printf("serial read error: %v", res.err)
			}

			out <- res.line
		}
	}
}

func ParseLoop(ctx context.Context, in <-chan string, out chan<- *pb.LocationRequest) {
	busvin := false

	for {
		select {
		case <-ctx.Done():
			return
		case sentence, ok := <-in:
			if !ok {
				return
			}

			s, err := nmea.ParseSentence(sentence)
			if err != nil {
				log.Printf("NMEA parse error: %v", err)
				continue
			}

			vin := util.If(busvin, "4S4BRDLC3B2413966", "4S4BRDLC3B2413967")

			loc, err := nmea.SentenceToLocation(s, vin)
			if err != nil {
				continue
			}
			out <- loc

			busvin = !busvin
		}
	}
}

func PublishLoop(ctx context.Context, producer *kafka.Producer, in <-chan *pb.LocationRequest) {
	topic := "vehicle-location"

	for {
		select {
		case <-ctx.Done():
			return
		case loc, ok := <-in:
			if !ok {
				return
			}

			if err := producer.PublishLocation(loc, &topic); err != nil {
				log.Printf("Kafka publish error: %v", err)
			}
		}
	}
}
