package main

import (
	"github.com/joho/godotenv"
	"github.com/phamphihungbk/go-graphql/configs"
	"github.com/phamphihungbk/go-graphql/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)


func main() {
	container := di.BuildContainer()

	err := container.Invoke(func(server *di.Server) {
		server.Run()
	})

	if err != nil {
		panic(err)
	}
}
