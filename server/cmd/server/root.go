package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/phamphihungbk/go-graphql/config"
	"log"
	"os"
)

type Application struct {
	Server *gin.Engine
}

func NewApplication(server *gin.Engine) *Application {
	return &Application{
		server,
	}
}

func Execute() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	dbCfg := config.NewDBCfg()
	e, err := InitializeApp(dbCfg)
	if err != nil {
		fmt.Printf("Cannot start app: %+v\n", err)
		os.Exit(1)
	}
	e.Server.Run(":8080")
}
