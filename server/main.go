package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/phamphihungbk/go-graphql/configs"
)

type App struct {
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	dbConfig := configs.NewDBCfg()
	if err := InitializeApp(dbConfig.GetConnectionInfo()); err != nil {
		fmt.Printf("Cannot start app: %+v\n", err)
		os.Exit(1)
	}

	fmt.Println(dbConfig.GetConnectionInfo())
	//app.Start()
}
