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
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	config := configs.GetConfig()

	db, err := gorm.Open(postgres.Open(config.Postgres.GetConnectionInfo()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	db.AutoMigrate(&migrations.User{}, &migrations.Test{})
}
