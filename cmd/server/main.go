package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	pb "github.com/jojohimawan/intelligent-agent-system/api"
	kafka "github.com/jojohimawan/intelligent-agent-system/internal/kafka"

	"github.com/adrianmo/go-nmea"
	"go.bug.st/serial"
)

func main() {
	kafkaProducer, err := kafka.NewProducer(
		"10.10.10.203:9092",
		"http://10.10.10.203:8085",
		"vehicle-location",
	)
	if err != nil {
		log.Fatalf("Failed to connect to Kafka: %v", err)
	}
	defer kafkaProducer.Close()

	mode := &serial.Mode{
		BaudRate: 9600,
	}

	port, err := serial.Open("/dev/ttyACM0", mode)
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
						log.Fatal(err)
					}

					if s.DataType() == nmea.TypeRMC {
						m := s.(nmea.RMC)
						fmt.Printf("Time: %s\n", m.Time)
						fmt.Printf("Validity: %s\n", m.Validity)
						fmt.Printf("Latitude GPS: %s\n", nmea.FormatGPS(m.Latitude))
						fmt.Printf("Longitude GPS: %s\n", nmea.FormatGPS(m.Longitude))
						fmt.Printf("Date: %s\n", m.Date)

						floated_latitude, err := strconv.ParseFloat(nmea.FormatGPS(m.Latitude), 64)
						if err != nil {
							log.Fatalf("Failed to convert latitude to float64: %v", err)
						}

						floated_longitude, err := strconv.ParseFloat(nmea.FormatGPS(m.Longitude), 64)
						if err != nil {
							log.Fatalf("Failed to convert longitude to float64: %v", err)
						}

						var location *pb.LocationRequest = &pb.LocationRequest{
							Vin:       "4S4BRDLC3B2413966",
							Lat:       floated_latitude,
							Lon:       floated_longitude,
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
