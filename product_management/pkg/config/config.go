package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppEnv       string `envconfig:"APP_ENV"`
	Port         int    `envconfig:"PORT"`
	AllowOrigins string `envconfig:"ALLOW_ORIGIN"`
	JwtSecret    string `envconfig:"JWT_SECRET"`
	ClientId     string `envconfig:"GOOGLE_CLIENT_ID"`

	DB struct {
		Region   string `envconfig:"REGION"`
		Endpoint string `envconfig:"ENDPOINT"`
	}
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

	API struct {
		BookAPIUrl string `envconfig:"DATA_API_URL"`
	}

	GrpcService struct {
		ProductGrpcHost string `envconfig:"PRODUCT_GRPC_HOST"`
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
