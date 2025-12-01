package config

import (
	"log"
	"os"

	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

type Config struct {
	KafkaBrokerURL    string `env:"KAFKA_BROKER_URL,required=true"`
	SchemaRegistryURL string `env:"SCHEMA_REGISTRY_URL,required=true"`
	SerialPort        string `env:"SERIAL_PORT,required=true"`
	CanNetwork        string `env:"CAN_NETWORK, required=true"`
	CanNetworkAddress string `env:"CAN_NETWORK_ADDRESS, required=true"`

	Environment string `env:"APP_ENV,default=development"`

	Extras env.EnvSet
}

func Load() (*Config, error) {
	var cfg Config

	if os.Getenv("APP_ENV") == "development" || os.Getenv("APP_ENV") == "" {
		if err := godotenv.Load(); err != nil {
			log.Printf("Warning: No .env file found: %v", err)
		}
	}

	extras, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		log.Fatal("Failed to unmarshal environment variables: %w", err)
		return nil, err
	}
	cfg.Extras = extras

	return &cfg, nil
}
