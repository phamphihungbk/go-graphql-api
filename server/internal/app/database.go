package app

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewDBConnection(configs *configs.DBCfg) *gorm.DB {
	db, err := gorm.Open(postgres.Open(configs.DBCfg), &gorm.Config{})
	if err != nil {
		log.Fatalf("gorm open database connection error")
	}
	db.AutoMigrate(&models.User{})

	return db
}
