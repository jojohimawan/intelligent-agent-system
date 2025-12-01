package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jojohimawan/intelligent-agent-system/internal/config"
	"github.com/jojohimawan/intelligent-agent-system/internal/kafka"
	"github.com/jojohimawan/intelligent-agent-system/internal/mapper"
	"github.com/jojohimawan/intelligent-agent-system/internal/pipeline"
	"github.com/jojohimawan/intelligent-agent-system/internal/serial"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	mongoStore, err := mapper.NewMongoStore(ctx, mapper.MongoConfig{
		URI:        cfg.MongoDBURI,
		Database:   cfg.MongoDBDatabase,
		Collection: cfg.MongoDBCollection,
		Username:   cfg.MongoDBAuthUser,
		Password:   cfg.MongoDBAuthPassword,
		Timeout:    2 * time.Second,
	})
	if err != nil {
		log.Printf("[WARN]Main: MongoDB connection failed: %v", err)
	}

	mapService := mapper.NewService(mongoStore, "./data/mappings.json")
	if err := mapService.Init(ctx); err != nil {
		log.Fatalf("[ERR]Main: Failed to init mappings: %v", err)
	}

	kafkaProducer, err := kafka.NewProducer(
		cfg.KafkaBrokerURL,
		cfg.SchemaRegistryURL,
	)
	if err != nil {
		log.Fatalf("Failed to connect to Kafka: %v", err)
	}
	defer kafkaProducer.Close()

	sr, err := serial.Open(cfg.SerialPort, 9600)
	if err != nil {
		fmt.Printf("failed to open serial port %s: %v", cfg.SerialPort, err)
	}
	defer sr.Close()

	sv, err := serial.Connect(ctx, cfg.CanNetwork, cfg.CanNetworkAddress)
	if err != nil {
		fmt.Printf("failed to open serial port %s: %v", cfg.CanNetworkAddress, err)
	}
	defer sv.Close()

	if err := pipeline.Run(ctx, sr, sv, mapService, kafkaProducer); err != nil {
		log.Fatalf("pipeline error: %v", err)
	}
}
