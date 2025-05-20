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

	MessageBroker struct {
		OrderBrokerHost string `envconfig:"ORDER_BROKER_HOST"`
		OrderTopic      string `envconfig:"ORDER_TOPIC"`
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
