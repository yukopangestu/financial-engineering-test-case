package api

import (
	"financial-engineering-test-case/internal/routes"
	"log"

	"github.com/labstack/echo/v4"
)

func StartApp() {
	// Create Echo instance
	e := echo.New()

	// Setup routes
	routes.SetupRoutes(e)

	// Start server
	port := ":11230"
	log.Printf("Starting server on %s", port)
	if err := e.Start(port); err != nil {
		log.Fatal(err)
	}
}
