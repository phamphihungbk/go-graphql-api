package app

import (
	"log"

	"github.com/phamphihungbk/go-graphql-api/internal/config"
	"github.com/phamphihungbk/go-graphql-api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConnection struct {
	isBooted bool
	DB       *gorm.DB
}

type IDatabaseConnection interface {
	Boot()
}

func NewDatabaseConnection(config config.Server) *DatabaseConnection {
	db, err := gorm.Open(postgres.Open(config.Database.ConnectionString()), &gorm.Config{})
	if err != nil {
		log.Fatalf("can't establish database connection")
	}

	return &DatabaseConnection{isBooted: false, DB: db}
}

func (d *DatabaseConnection) Boot() {
	if d.isBooted {
		return
	}

	d.runMigration()
	d.isBooted = true
}

func (d *DatabaseConnection) runMigration() {
	d.DB.AutoMigrate(&model.User{})
}
