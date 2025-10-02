package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
	"time"

	pb "github.com/jojohimawan/intelligent-agent-system/api"
	"github.com/jojohimawan/intelligent-agent-system/internal/config"
	kafka "github.com/jojohimawan/intelligent-agent-system/internal/kafka"

	"github.com/adrianmo/go-nmea"
	"go.bug.st/serial"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	kafkaProducer, err := kafka.NewProducer(
		cfg.KafkaBrokerURL,
		cfg.SchemaRegistryURL,
		cfg.KafkaTopicVehicleLocation,
	)
	if err != nil {
		log.Fatalf("Failed to connect to Kafka: %v", err)
	}
	defer kafkaProducer.Close()

	mode := &serial.Mode{
		BaudRate: 9600,
	}

	port, err := serial.Open(cfg.SerialPort, mode)
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	reader := bufio.NewReader(port)
	var buffer string

	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
			continue
		}

		buffer += data

		if strings.Contains(buffer, "\n") {
			sentences := strings.Split(buffer, "\n")

			for i := 0; i < len(sentences)-1; i++ {
				sentence := strings.TrimSpace(sentences[i])

				if sentence != "" {
					fmt.Printf("%s", sentence)
					fmt.Printf("\n")

					s, err := nmea.Parse(sentence)
					if err != nil {
						log.Printf("NMEA parse error: %v (sentence: %s)", err, sentence)
						continue
					}

					if s.DataType() == nmea.TypeRMC {
						m := s.(nmea.RMC)
						fmt.Printf("Time: %s\n", m.Time)
						fmt.Printf("Validity: %s\n", m.Validity)
						fmt.Printf("Latitude GPS: %f\n", m.Latitude)
						fmt.Printf("Longitude GPS: %f\n", m.Longitude)
						fmt.Printf("Date: %s\n", m.Date)

						var location *pb.LocationRequest = &pb.LocationRequest{
							Vin:       "4S4BRDLC3B2413966",
							Lat:       m.Latitude,
							Lon:       m.Longitude,
							Timestamp: time.Now().Unix(),
						}

						if err := kafkaProducer.PublishLocation(location); err != nil {
							log.Printf("Kafka publish error: %v", err)
						}
					}
				}
			}

			buffer = sentences[len(sentences)-1]
		}
	}
}
