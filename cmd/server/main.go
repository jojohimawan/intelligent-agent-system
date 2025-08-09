package main

import (
	"context"
	"log"
	"net"

	pb "github.com/jojohimawan/intelligent-agent-system/api"
    kafka "github.com/jojohimawan/intelligent-agent-system/internal/kafka"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedLocationServiceServer
    producer *kafka.Producer
}

func (s *server) SendLocation(ctx context.Context, req *pb.LocationRequest) (*pb.LocationResponse, error) {
    log.Printf("Received location: lat=%.6f, lon=%.6f, ts=%d", req.Lat, req.Lon, req.Timestamp)

	err := s.producer.PublishLocation(ctx, req)
	if err != nil {
		return &pb.LocationResponse{
        Lat: req.Lat,
        Lon: req.Lon,
        Timestamp: req.Timestamp,
        Success: false,
        Status: "kafka_error",
   		}, nil
	}

    return &pb.LocationResponse{
        Lat:       req.Lat,
        Lon:       req.Lon,
        Timestamp: req.Timestamp,
        Success:   true,
        Status:    "processed",
    }, nil
}

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

    grpcServer := grpc.NewServer()
    pb.RegisterLocationServiceServer(grpcServer, &server{
        producer: kafkaProducer,
    })


    log.Println("gRPC server running on port 50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}