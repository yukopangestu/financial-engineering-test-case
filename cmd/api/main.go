package api

import (
	"financial-engineering-test-case/internal/config"
	"financial-engineering-test-case/internal/database"
	"financial-engineering-test-case/internal/routes"
	"log"

	"github.com/labstack/echo/v4"
)

func StartApp() {
	// Load configuration
	cfg := config.LoadConfig()
	log.Println("Configuration loaded")

	// Initialize database
	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run auto-migration
	if err := database.AutoMigrate(); err != nil {
		log.Fatalf("Failed to run auto-migration: %v", err)
	}

	// Create Echo instance
	e := echo.New()

	// Setup routes
	routes.SetupRoutes(e, db)

	// Start server
	port := ":11230"
	log.Printf("Starting server on %s", port)
	if err := e.Start(port); err != nil {
		log.Fatal(err)
	}
}
