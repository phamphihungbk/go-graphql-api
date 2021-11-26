package config

import (
	"os"
)

type AppCfg struct {
	Env          string `env:"APP_ENV"`
	DBConnection string `env:"DB_CONNECTION"`
	Host         string `env:"APP_HOST"`
	Port         string `env:"APP_PORT"`
}

func NewAppCfg() *AppCfg {
	return &AppCfg{
		Env:          os.Getenv("APP_ENV"),
		DBConnection: os.Getenv("DB_CONNECTION"),
		Host:         os.Getenv("APP_HOST"),
		Port:         os.Getenv("APP_PORT"),
	}
}
