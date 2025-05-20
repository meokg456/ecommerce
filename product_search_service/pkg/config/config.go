package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppEnv       string `envconfig:"APP_ENV"`
	Port         int    `envconfig:"PORT"`
	AllowOrigins string `envconfig:"ALLOW_ORIGIN"`

	PostgresDB struct {
		Host      string `envconfig:"DB_HOST"`
		User      string `envconfig:"DB_USER"`
		Password  string `envconfig:"DB_PASS"`
		DBName    string `envconfig:"DB_NAME"`
		Port      string `envconfig:"DB_PORT"`
		EnableSSL bool   `envconfig:"ENABLE_SSL"`
	}

	AwsService struct {
		Region   string `envconfig:"AWS_REGION"`
		Endpoint string `envconfig:"AWS_ENDPOINT"`
		Bucket   string `envconfig:"AWS_S3_BUCKET"`
	}

	MessageBroker struct {
		ProductBrokerHost string `envconfig:"PRODUCT_BROKER_HOST"`
		ProductTopic      string `envconfig:"PRODUCT_TOPIC"`
		ProductGroupId    string `envconfig:"PRODUCT_GROUP_ID"`
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
