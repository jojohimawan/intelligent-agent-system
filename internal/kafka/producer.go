package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	srclient "github.com/riferrei/srclient"

	pb "github.com/jojohimawan/intelligent-agent-system/api"
)

type Producer struct {
	kafkaProducer *ckafka.Producer
	schemaRegistryClient *srclient.SchemaRegistryClient
	topic string
	schemaID int
}

func NewProducer(broker, schemaRegistryURL, topic, jsonSchemaPath string) (*Producer, error) {
	p, err := ckafka.NewProducer(&ckafka.ConfigMap{
		"bootstrap.servers": broker,
	})
	if err != nil {
		return nil, fmt.Errorf("Failed to create producer: %w", err)
	}

	src := srclient.NewSchemaRegistryClient(schemaRegistryURL)

	schemaBytes, err := os.ReadFile(jsonSchemaPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to read JSON schema: %w", err)
	}

	subject := topic + "-value"

	schema, err := src.GetLatestSchema(subject)
    if err != nil {
        log.Printf("Schema not found for subject %s, registering new...", subject)
        schema, err = src.CreateSchema(subject, string(schemaBytes), srclient.Json)
        if err != nil {
            return nil, fmt.Errorf("failed to register schema: %w", err)
        }
        log.Printf("Registered schema ID %d for subject %s", schema.ID(), subject)
    } else {
        log.Printf("Found existing schema ID %d for subject %s", schema.ID(), subject)
    }

	return &Producer{
		kafkaProducer: p,
		schemaRegistryClient: src,
		topic: topic,
		schemaID: schema.ID(),
	}, nil
}

func (p *Producer) PublishLocation(ctx context.Context, loc *pb.LocationRequest) error {
	payload, err := json.Marshal(loc)
	if err != nil {
		return fmt.Errorf("Failed to marshal JSON: %w", err)
	}

	msgBytes := make([]byte, 5+len(payload))
	msgBytes[0] = 0
    msgBytes[1] = byte(p.schemaID >> 24)
    msgBytes[2] = byte(p.schemaID >> 16)
    msgBytes[3] = byte(p.schemaID >> 8)
    msgBytes[4] = byte(p.schemaID)
    copy(msgBytes[5:], payload)

	return p.kafkaProducer.Produce(&ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &p.topic, Partition: ckafka.PartitionAny},
		Value: msgBytes,
	}, nil)
}

func (p *Producer) Close() {
	p.kafkaProducer.Close()
}