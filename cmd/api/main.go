package api

import (
	"financial-engineering-test-case/internal/config"
	"financial-engineering-test-case/internal/database"
	"financial-engineering-test-case/internal/routes"
	"log"

	"github.com/labstack/echo/v4"
)

func StartApp() {
	cfg := config.LoadConfig()
	log.Println("Configuration loaded")

	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := database.AutoMigrate(); err != nil {
		log.Fatalf("Failed to run auto-migration: %v", err)
	}

	e := echo.New()

	routes.SetupRoutes(e, db, cfg)

	port := ":11230"
	log.Printf("Starting server on %s", port)
	if err := e.Start(port); err != nil {
		log.Fatal(err)
	}
}
