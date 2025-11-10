package main

import (
	"context"
	"fmt"
	"log"

	bit29can "github.com/jojohimawan/intelligent-agent-system/internal/can/gen"
	"github.com/jojohimawan/intelligent-agent-system/internal/config"
	"github.com/jojohimawan/intelligent-agent-system/internal/kafka"
	"github.com/jojohimawan/intelligent-agent-system/internal/pipeline"
	"github.com/jojohimawan/intelligent-agent-system/internal/serial"
)

func main() {
	ctx := context.Background()

	defineSignals()
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	kafkaProducer, err := kafka.NewProducer(
		cfg.KafkaBrokerURL,
		cfg.SchemaRegistryURL,
		"vehicle-obd",
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

	sv, err := serial.Connect(ctx, "can", "vcan0")
	if err != nil {
		fmt.Printf("failed to open serial port %s: %v", "vcan0", err)
	}
	defer sv.Close()

	if err := pipeline.Run(ctx, sr, sv, kafkaProducer); err != nil {
		log.Fatalf("pipeline error: %v", err)
	}
}

func defineSignals() {
	md := bit29can.Messages().OBD2
	fmt.Println("=== OBD2 Signal Definitions ===")

	fmt.Printf("Length Signal:\n")
	fmt.Printf("  Start: %d, Length: %d, BigEndian: %v\n\n",
		md.Length.Start, md.Length.Length, md.Length.IsBigEndian)

	fmt.Printf("Service Signal:\n")
	fmt.Printf("  Start: %d, Length: %d, BigEndian: %v\n\n",
		md.Service.Start, md.Service.Length, md.Service.IsBigEndian)

	fmt.Printf("S01PID Signal:\n")
	fmt.Printf("  Start: %d, Length: %d, BigEndian: %v\n\n",
		md.S01PID.Start, md.S01PID.Length, md.S01PID.IsBigEndian)

	fmt.Printf("S01PID0C_EngineRPM Signal:\n")
	fmt.Printf("  Start: %d, Length: %d, BigEndian: %v\n",
		md.S01PID0C_EngineRPM.Start, md.S01PID0C_EngineRPM.Length,
		md.S01PID0C_EngineRPM.IsBigEndian)
	fmt.Printf("  Scale: %f, Offset: %f\n",
		md.S01PID0C_EngineRPM.Scale, md.S01PID0C_EngineRPM.Offset)
}
