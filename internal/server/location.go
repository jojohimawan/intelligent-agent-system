package server

import (
	"log"
	"context"

	pb "github.com/jojohimawan/intelligent-agent-system/api"
	kafka "github.com/jojohimawan/intelligent-agent-system/internal/kafka"
)

type LocationServer struct {
	pb.UnimplementedLocationServiceServer
    producer *kafka.Producer
}

func NewLocationServer(producer *kafka.Producer) *LocationServer {
	return &LocationServer{producer: producer}
}

func (s *LocationServer) SendLocation(ctx context.Context, req *pb.LocationRequest) (*pb.LocationResponse, error) {
    log.Printf("Received location: vin=%s, lat=%.6f, lon=%.6f, ts=%d", req.Vin, req.Lat, req.Lon, req.Timestamp)

	err := s.producer.PublishLocation(ctx, req)
	if err != nil {
		return &pb.LocationResponse{
        Vin: req.Vin,
        Lat: req.Lat,
        Lon: req.Lon,
        Timestamp: req.Timestamp,
        Success: false,
        Status: "kafka_error",
   		}, nil
	}

    return &pb.LocationResponse{
        Vin: req.Vin,
        Lat:       req.Lat,
        Lon:       req.Lon,
        Timestamp: req.Timestamp,
        Success:   true,
        Status:    "processed",
    }, nil
}