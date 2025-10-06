package main

import (
	"context"
	"log"

	"github.com/jojohimawan/intelligent-agent-system/internal/config"
	"github.com/jojohimawan/intelligent-agent-system/internal/kafka"
	"github.com/jojohimawan/intelligent-agent-system/internal/pipeline"
	"github.com/jojohimawan/intelligent-agent-system/internal/serial"
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

	sr, err := serial.Open(cfg.SerialPort, 9600)
	if err != nil {
		log.Fatalf("failed to open serial port %s: %v", cfg.SerialPort, err)
	}
	defer sr.Close()

	ctx := context.Background()
	if err := pipeline.Run(ctx, sr, kafkaProducer); err != nil {
		log.Fatalf("pipeline error: %v", err)
	}
}
