package configs

import (
	"os"
)

// Config object
type Config struct {
	Env          string         `env:"APP_ENV"`
	dbConnection string         `env:"DB_CONNECTION"`
	Postgres     PostgresConfig `json:"postgres"`
}

func GetConfig() Config {
	return Config{
		Env:          os.Getenv("APP_ENV"),
		dbConnection: os.Getenv("DB_CONNECTION"),
		Postgres:     GetPostgresConfig(),
	}
}
