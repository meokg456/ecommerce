package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppEnv       string `envconfig:"APP_ENV"`
	Port         int    `envconfig:"PORT"`
	GrpcPort     int    `envconfig:"GRPC_PORT"`
	AllowOrigins string `envconfig:"ALLOW_ORIGIN"`

	DB struct {
		Region   string `envconfig:"AWS_REGION"`
		Endpoint string `envconfig:"AWS_ENDPOINT"`
	}

	MessageBroker struct {
		ProductBrokerHost string `envconfig:"PRODUCT_BROKER_HOST"`
		ProductTopic      string `envconfig:"PRODUCT_TOPIC"`
	}
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	config := new(Config)

	err = envconfig.Process("", config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
