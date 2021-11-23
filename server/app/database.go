package database

import (
	"github.com/phamphihungbk/go-graphql/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"github.com/phamphihungbk/go-graphql/configs"
)

func NewDBConnection(configs *configs.DBCfg) *gorm.DB {
	db, err := gorm.Open(postgres.Open(configs.DBCfg), &gorm.Config{})
	if err != nil {
		log.Fatalf("gorm open database connection error")
	}
	db.AutoMigrate(&models.User{})

	return db
}
