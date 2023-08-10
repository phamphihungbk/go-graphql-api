package config

import (
	"github.com/phamphihungbk/go-graphql-api/internal/util"
)

type Server struct {
	Env        string `env:"APP_ENV"`
	Connection string `env:"DB_CONNECTION"`
	Host       string `env:"APP_HOST"`
	Port       string `env:"APP_PORT"`
	Database   Database
}

func DefaultConfigFromEnv() Server {
	return Server{
		Env:        util.GetEnv("APP_ENV", "dev"),
		Connection: util.GetEnv("DB_CONNECTION", "postgres"),
		Host:       util.GetEnv("APP_HOST", "localhost"),
		Port:       util.GetEnv("APP_PORT", "8080"),
		Database: Database{
			Host:         util.GetEnv("DB_HOST", "graphql-db"),
			Port:         util.GetEnvAsInt("DB_PORT", 5432),
			User:         util.GetEnv("DB_USER", "graphql_user"),
			Password:     util.GetEnv("DB_PASSWORD", "graphql_pass"),
			DatabaseName: util.GetEnv("DB_NAME", "graphql_db"),
		},
	}
}
