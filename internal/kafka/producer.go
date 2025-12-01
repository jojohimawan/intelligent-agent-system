package kafka

import (
	"context"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde/protobuf"

	pb "github.com/jojohimawan/intelligent-agent-system/api"
)

type Producer struct {
	kafkaProducer        *ckafka.Producer
	schemaRegistryClient schemaregistry.Client
	protobufSerde        *protobuf.Serializer
}

func NewProducer(broker, schemaRegistryURL string) (*Producer, error) {
	p, err := ckafka.NewProducer(&ckafka.ConfigMap{
		"bootstrap.servers": broker,
		"client.id":         "ias-go-producer",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}

	src, err := schemaregistry.NewClient(
		schemaregistry.NewConfig(schemaRegistryURL),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create schema registry client: %w", err)
	}

	serdeConfig := protobuf.NewSerializerConfig()
	serdeConfig.AutoRegisterSchemas = true
	serdeConfig.UseLatestVersion = true

	serializer, err := protobuf.NewSerializer(src, serde.ValueSerde, serdeConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create protobuf serializer: %w", err)
	}

	return &Producer{
		kafkaProducer:        p,
		schemaRegistryClient: src,
		protobufSerde:        serializer,
	}, nil
}

func (p *Producer) PublishLocation(loc *pb.LocationRequest, topic *string, option ...context.Context) error {
	serializedPayload, err := p.protobufSerde.Serialize(*topic, loc)
	if err != nil {
		return fmt.Errorf("failed to serialize protobuf message: %w", err)
	}

	return p.kafkaProducer.Produce(&ckafka.Message{
		TopicPartition: ckafka.TopicPartition{
			Topic:     topic,
			Partition: ckafka.PartitionAny,
		},
		Value: serializedPayload,
		Headers: []ckafka.Header{
			{
				Key:   "content-type",
				Value: []byte("application/x-protobuf"),
			},
		},
	}, nil)
}

func (p *Producer) PublishTelematics(telematics *pb.TelematicsBatch, topic *string, option ...context.Context) error {
	serializedPayload, err := p.protobufSerde.Serialize(*topic, telematics)
	if err != nil {
		return fmt.Errorf("failed to serialize protobuf message: %w", err)
	}

	return p.kafkaProducer.Produce(&ckafka.Message{
		TopicPartition: ckafka.TopicPartition{
			Topic:     topic,
			Partition: ckafka.PartitionAny,
		},
		Key:   []byte(telematics.Vin),
		Value: serializedPayload,
		Headers: []ckafka.Header{
			{
				Key:   "content-type",
				Value: []byte("application/x-protobuf"),
			},
		},
	}, nil)
}

func (p *Producer) Close() {
	p.kafkaProducer.Close()
}

func (p *Producer) Flush() {
	p.kafkaProducer.Flush(15 * 1000)
}
