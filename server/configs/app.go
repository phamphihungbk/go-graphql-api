package configs

import (
	"os"
)

type AppCfg struct {
	env          string `env:"APP_ENV"`
	dbConnection string `env:"DB_CONNECTION"`
	host         string `env:"APP_HOST"`
	port         string `env:"APP_PORT"`
}

func NewAppCfg() *AppCfg {
	return &AppCfg{
		env:          os.Getenv("APP_ENV"),
		dbConnection: os.Getenv("DB_CONNECTION"),
		host:         os.Getenv("APP_HOST"),
		port:         os.Getenv("APP_PORT"),
	}
}
