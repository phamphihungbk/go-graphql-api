package config

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
	DBCfg    string
}

func NewDBCfg() *DBCfg {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err)
	}
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbCfg := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	return &DBCfg{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		dbName:  dbName,
		DBCfg: dbCfg,
	}
}
