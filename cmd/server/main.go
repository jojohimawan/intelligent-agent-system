package main

import (
	"context"
	"log"
	"net"
    "time"

	pb "github.com/jojohimawan/intelligent-agent-system/api"
    kafka "github.com/jojohimawan/intelligent-agent-system/internal/kafka"
    grpcServer "github.com/jojohimawan/intelligent-agent-system/internal/server"
    "github.com/jojohimawan/intelligent-agent-system/internal/data"

	"google.golang.org/grpc"
)

func main() {
	kafkaProducer, err := kafka.NewProducer(
		"10.10.10.203:9092",
		"http://10.10.10.203:8085",
		"vehicle-location",
		"./schema/location_schema.json",
	)
	if err != nil {
		log.Fatalf("Failed to connect to Kafka: %v", err)
	}
	defer kafkaProducer.Close()

    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcSrv := grpc.NewServer()
    pb.RegisterLocationServiceServer(grpcSrv, grpcServer.NewLocationServer(kafkaProducer))

    locations, err := data.LoadDummyLocations("dummy/vehicle_locations.csv")
    if err != nil {
        log.Fatalf("Failed to load CSV: %v", err)
    }

    go func() {
        ctx := context.Background()

        for index, loc := range locations {
            log.Printf("Publishing Row %d of VIN=%s to Kafka. lat=%.6f, lon=%.6f, ts=%d", index, loc.Vin, loc.Lat, loc.Lon, loc.Timestamp,)
            if err := kafkaProducer.PublishLocation(ctx, loc); err != nil {
                log.Printf("Kafka publish error: %v", err)
            }

            time.Sleep(3 * time.Second)
        }
    }()

    log.Println("gRPC server running on port 50051")
    if err := grpcSrv.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}