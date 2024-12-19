package database

import (
	"log"
	configuration "secret-satana/configs"
	"secret-satana/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error

	file := "app.db"
	if configuration.DBFileLocation != "" {
		file = configuration.DBFileLocation
	}

	DB, err = gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate models
	err = DB.AutoMigrate(&models.User{}, &models.SatanaSelection{})
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Database initialized and migrations applied!")
}
