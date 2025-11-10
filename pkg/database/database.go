package database

import (
	"fmt"
	"log"

	"go-gin-starter/config"
	"go-gin-starter/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the database connection
var DB *gorm.DB

// ConnectDB connects to the database
func ConnectDB(cfg *config.Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established")
}

// AutoMigrate migrates the database
func AutoMigrate() {
	err := DB.AutoMigrate(&models.User{}, &models.Book{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration completed")
}
