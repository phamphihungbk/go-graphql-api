package server

import (
	"fmt"
	"log"
	"os"

	"github.com/phamphihungbk/go-graphql-api/internal/app"

	"github.com/joho/godotenv"
	"github.com/phamphihungbk/go-graphql-api/internal/config"
)

type App struct {
	router     *app.Router
	connection *app.DatabaseConnection
}

func NewApp(router *app.Router, connection *app.DatabaseConnection) *App {
	return &App{
		router:     router,
		connection: connection,
	}
}

func Execute() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	config := config.DefaultConfigFromEnv()
	app, err := InitializeApp(config)

	if err != nil {
		fmt.Printf("Cannot start app: %+v\n", err)
		os.Exit(1)
	}

	app.connection.Boot()
	app.router.Boot()
}
