package configs

import (
	"os"
)

type Config struct {
	Env          string   `env:"APP_ENV"`
	dbConnection string   `env:"DB_CONNECTION"`
	Host         string   `env:"APP_HOST"`
	Port         string   `env:"APP_PORT"`
	Postgres     DBConfig `json:"postgres"`
}

func GetConfig() Config {
	return Config{
		Env:          os.Getenv("APP_ENV"),
		dbConnection: os.Getenv("DB_CONNECTION"),
		Host:         os.Getenv("APP_HOST"),
		Port:         os.Getenv("APP_PORT"),
		Postgres:     GetDBConfig(),
	}
}
