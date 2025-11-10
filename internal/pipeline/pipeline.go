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

	pb "github.com/jojohimawan/intelligent-agent-system/api"
	internalcan "github.com/jojohimawan/intelligent-agent-system/internal/can"
	bit29can "github.com/jojohimawan/intelligent-agent-system/internal/can/gen"
	"github.com/jojohimawan/intelligent-agent-system/internal/kafka"
	"github.com/jojohimawan/intelligent-agent-system/internal/nmea"
	"github.com/jojohimawan/intelligent-agent-system/internal/serial"
	"github.com/jojohimawan/intelligent-agent-system/internal/util"
	"go.einride.tech/can"
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
	decodedFrame := make(chan *internalcan.OBD2, 50)
	parsedFrame := make(chan *pb.VehicleOBD, 50)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

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
		DecodeFrameLoop(ctx, rawFrame, decodedFrame)
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
	for {
		select {
		case <-ctx.Done():
			return
		default:
			for sr.Recv.Receive() {
				frame := sr.Recv.Frame()

				if frame.ID != bit29can.Messages().OBD2.ID || !frame.IsExtended {
					continue
				}

				out <- frame
			}
		}
	}
}

func DecodeFrameLoop(ctx context.Context, in <-chan can.Frame, out chan<- *internalcan.OBD2) {
	for {
		select {
		case <-ctx.Done():
			return
		case frame := <-in:
			decodedFrame, err := internalcan.DecodeMode01PID(frame)
			if err != nil {
				fmt.Println("%v", err)
				continue
			}

			out <- decodedFrame
		}
	}
}

func ParseFrameLoop(ctx context.Context, in <-chan *internalcan.OBD2, out chan<- *pb.VehicleOBD) {
	for {
		select {
		case <-ctx.Done():
			return
		case decodedFrame := <-in:
			parsedMsg, err := internalcan.MessageToOBD("4S4BRDLC3B2413966", decodedFrame)
			if err != nil {
				fmt.Println("%v", err)
				continue
			}

			out <- parsedMsg
		}
	}
}

func PublishFrameLoop(ctx context.Context, producer *kafka.Producer, in <-chan *pb.VehicleOBD) {
	topic := "vehicle-obd"

	for {
		select {
		case <-ctx.Done():
			return
		case parsedMsg := <-in:
			if err := producer.PublishOBD(parsedMsg, &topic); err != nil {
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

	var buffer string

	for {
		select {
		case <-ctx.Done():
			return
		default:
			line, err := sr.ReadLine()
			if err != nil {
				log.Printf("serial read error: %v", err)
				continue
			}

			buffer += line

			if strings.Contains(buffer, "\n") {
				parts := strings.Split(buffer, "\n")

				for i := 0; i < len(parts)-1; i++ {
					s := strings.TrimSpace(parts[i])
					if s != "" {
						out <- s
					}
				}

				buffer = parts[len(parts)-1]
			}
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
