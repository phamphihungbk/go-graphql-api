package configs

import (
	"fmt"
	"os"
	"strconv"
)

type DBCfg struct {
	host     string `env:"DB_HOST"`
	port     int    `env:"DB_PORT"`
	user     string `env:"DB_USER"`
	password string `env:"DB_PASSWORD"`
	dbName   string `env:"DB_NAME"`
}

func (c *DBCfg) GetConnectionInfo() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbName=%s sslmode=disable",
		c.host, c.port, c.user, c.password, c.dbName)
}

func NewDBCfg() *DBCfg {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err)
	}

	return &DBCfg{
		host:     os.Getenv("DB_HOST"),
		port:     port,
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		dbName:   os.Getenv("DB_NAME"),
	}
}
