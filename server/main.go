package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/phamphihungbk/go-graphql/configs"
	"log"
	"os"
	"github.com/gin-gonic/gin"
)

type Application struct {
	Router *gin.Engine
}

func NewApplication(router *gin.Engine) *Application {
	return &Application{
		Router: router,
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	dbCfg := configs.NewDBCfg()
	e, err := InitializeApp(dbCfg)
	if err != nil {
		fmt.Printf("Cannot start app: %+v\n", err)
		os.Exit(1)
	}
	e.Router.Run(":8080")
}
