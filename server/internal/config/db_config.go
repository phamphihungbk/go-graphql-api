package config

import (
	"fmt"
)

type Database struct {
	Host         string `env:"DB_HOST"`
	Port         int    `env:"DB_PORT"`
	User         string `env:"DB_USER"`
	Password     string `env:"DB_PASSWORD"`
	DatabaseName string `env:"DB_NAME"`
}

func (c Database) ConnectionString() string {
	return fmt.Sprintf(
		"Host=%s Port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.DatabaseName,
	)
}
