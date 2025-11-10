package main

import (
	"log"

	"go-gin-starter/config"
	"go-gin-starter/internal/routes"
	"go-gin-starter/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Connect to the database
	database.ConnectDB(cfg)

	// Migrate the database
	database.AutoMigrate()

	// Initialize Gin router
	router := gin.Default()

	// Setup routes
	routes.AuthRoutes(router)
	routes.BookRoutes(router)

	// Start the server
	port := cfg.APIPort
	if port == "" {
		port = "8080" // Default port
	}
	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
