package main

import (
	"context"
	"log"
	"net"

	pb "github.com/jojohimawan/intelligent-agent-system.git/api"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedLocationServiceServer
}

func (s *server) SendLocation(ctx context.Context, req *pb.LocationRequest) (*pb.LocationResponse, error) {
    log.Printf("Received location: lat=%.6f, lon=%.6f, ts=%d", req.Lat, req.Lon, req.Timestamp)
    return &pb.LocationResponse{
        Lat:       req.Lat,
        Lon:       req.Lon,
        Timestamp: req.Timestamp,
        Success:   true,
        Status:    "processed",
    }, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterLocationServiceServer(grpcServer, &server{})

    log.Println("gRPC server running on port 50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}