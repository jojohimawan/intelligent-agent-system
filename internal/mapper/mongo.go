package mapper

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoConfig struct {
	URI        string
	Database   string
	Collection string
	Username   string
	Password   string
	Timeout    time.Duration
}

type MongoStore struct {
	client     *mongo.Client
	collection *mongo.Collection
	config     MongoConfig
}

func NewMongoStore(ctx context.Context, cfg MongoConfig) (*MongoStore, error) {
	clientOptions := options.Client().ApplyURI(cfg.URI)
	if cfg.Username != "" || cfg.Password != "" {
		clientOptions.SetAuth(options.Credential{
			Username:   cfg.Username,
			Password:   cfg.Password,
			AuthSource: cfg.Database,
		})
	} else {
		return nil, fmt.Errorf("no mongo credentials provided")
	}

	connectCtx, cancel := context.WithTimeout(ctx, cfg.Timeout)
	defer cancel()

	client, err := mongo.Connect(connectCtx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to create mongo client: %w", err)
	}

	if err := client.Ping(connectCtx, readpref.Primary()); err != nil {
		_ = client.Disconnect(context.Background())
		return nil, fmt.Errorf("mongo network unreachable: %w", err)
	}

	return &MongoStore{
		client:     client,
		collection: client.Database(cfg.Database).Collection(cfg.Collection),
		config:     cfg,
	}, nil
}

func (m *MongoStore) Close(ctx context.Context) error {
	if m.client != nil {
		return m.client.Disconnect(ctx)
	}

	return nil
}
