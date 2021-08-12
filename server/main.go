package main

import (
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	"log"
	"github.com/phamphihungbk/go-graphql/configs"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := configs.GetConfig()
	println(config.Postgres.GetPostgresConnectionInfo())
	Connects to PostgresDB
	db, err := gorm.Open(
		config.Postgres.Dialect(),
		config.Postgres.GetPostgresConnectionInfo(),
	)

	if err != nil {
		panic(err)
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})
}