package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/phamphihungbk/go-graphql/config"
	"log"
	"os"
	"github.com/gin-gonic/gin"
)

type Application struct {
	AppCfg *config.AppCfg
	Router *gin.Engine
}

func NewApplication(appCfg *config.AppCfg, router *gin.Engine) *Application {
	return &Application{
		AppCfg: appCfg,
		Router: router,
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	appCfg := config.NewAppCfg()
	dbCfg := config.NewDBCfg()
	e, err := InitializeApp(appCfg, dbCfg)
	if err != nil {
		fmt.Printf("Cannot start app: %+v\n", err)
		os.Exit(1)
	}
	e.Router.Run(":8080")
}
