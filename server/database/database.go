package database

import (
	"github.com/phamphihungbk/go-graphql/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewDBConnection(DBCfg string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(DBCfg), &gorm.Config{})
	if err != nil {
		log.Fatalf("gorm open database connection error")
	}
	db.AutoMigrate(&models.User{})

	return db
}
