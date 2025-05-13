package postgresstore

import (
	"fmt"

	"github.com/meokg456/sampleservice/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Options struct {
	Host      string
	User      string
	Password  string
	DBName    string
	Port      string
	EnableSSL bool
}

func ParseFromConfig(c *config.Config) Options {
	return Options{
		Host:      c.PostgresDB.Host,
		User:      c.PostgresDB.User,
		Password:  c.PostgresDB.Password,
		DBName:    c.PostgresDB.DBName,
		Port:      c.PostgresDB.Port,
		EnableSSL: c.PostgresDB.EnableSSL,
	}
}

func NewConnection(options Options) (*gorm.DB, error) {
	sslmode := "disable"

	if options.EnableSSL {
		sslmode = "enable"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", options.Host, options.User, options.Password, options.DBName, options.Port, sslmode)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
