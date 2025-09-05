package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"

	// "context"
	// "net"
	// "time"

	// pb "github.com/jojohimawan/intelligent-agent-system/api"
	// kafka "github.com/jojohimawan/intelligent-agent-system/internal/kafka"
	// grpcServer "github.com/jojohimawan/intelligent-agent-system/internal/server"
	// "github.com/jojohimawan/intelligent-agent-system/internal/data"

	// "google.golang.org/grpc"

	"github.com/adrianmo/go-nmea"
	"go.bug.st/serial"
)

func main() {
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
					fmt.Printf(sentence)
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
					}
				}
			}

			buffer = sentences[len(sentences)-1]
		}
	}
	// kafkaProducer, err := kafka.NewProducer(
	// 	"10.10.10.203:9092",
	// 	"http://10.10.10.203:8085",
	// 	"vehicle-location",
	// 	"./schema/location_schema.json",
	// )
	// if err != nil {
	// 	log.Fatalf("Failed to connect to Kafka: %v", err)
	// }
	// defer kafkaProducer.Close()

	// lis, err := net.Listen("tcp", ":50051")
	// if err != nil {
	//     log.Fatalf("Failed to listen: %v", err)
	// }

	// grpcSrv := grpc.NewServer()
	// pb.RegisterLocationServiceServer(grpcSrv, grpcServer.NewLocationServer(kafkaProducer))

	// locations, err := data.LoadDummyLocations("dummy/vehicle_locations.csv")
	// if err != nil {
	//     log.Fatalf("Failed to load CSV: %v", err)
	// }

	// go func() {
	//     ctx := context.Background()

	//     for index, loc := range locations {
	//         log.Printf("Publishing Row %d of VIN=%s to Kafka. lat=%.6f, lon=%.6f, ts=%d", index, loc.Vin, loc.Lat, loc.Lon, loc.Timestamp,)
	//         if err := kafkaProducer.PublishLocation(ctx, loc); err != nil {
	//             log.Printf("Kafka publish error: %v", err)
	//         }

	//         time.Sleep(3 * time.Second)
	//     }
	// }()

	// log.Println("gRPC server running on port 50051")
	// if err := grpcSrv.Serve(lis); err != nil {
	//     log.Fatalf("Failed to serve: %v", err)
	// }
}
