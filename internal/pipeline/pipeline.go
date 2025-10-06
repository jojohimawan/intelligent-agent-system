package pipeline

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	pb "github.com/jojohimawan/intelligent-agent-system/api"
	"github.com/jojohimawan/intelligent-agent-system/internal/kafka"
	"github.com/jojohimawan/intelligent-agent-system/internal/nmea"
	"github.com/jojohimawan/intelligent-agent-system/internal/serial"
	"github.com/jojohimawan/intelligent-agent-system/internal/util"
)

func Run(
	ctx context.Context,
	sr *serial.SerialReader,
	producer *kafka.Producer,
) error {
	rawSentences := make(chan string, 50)
	locations := make(chan *pb.LocationRequest, 50)

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

func ReadSerialLoop(ctx context.Context, sr *serial.SerialReader, out chan<- string) {
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
	for {
		select {
		case <-ctx.Done():
			return
		case loc, ok := <-in:
			if !ok {
				return
			}

			if err := producer.PublishLocation(loc); err != nil {
				log.Printf("Kafka publish error: %v", err)
			}
		}
	}
}
