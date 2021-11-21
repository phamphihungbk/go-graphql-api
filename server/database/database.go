package database

import (
	"github.com/phamphihungbk/go-graphql/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Manager struct {
	DB *gorm.DB
}

func NewDBConnection(connectionInfo string) *Manager {
	db, err := gorm.Open(postgres.Open(connectionInfo), &gorm.Config{})
	if err != nil {
		log.Fatalf("gorm open database connection error")
	}
	db.AutoMigrate(&models.User{})

	return &Manager{db}
}
