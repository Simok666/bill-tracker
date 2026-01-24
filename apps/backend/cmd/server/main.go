package main

import (
	"log"

	"github.com/dhani/bill-tracker-backend/internal/config"
	"github.com/dhani/bill-tracker-backend/internal/database"
	"github.com/dhani/bill-tracker-backend/internal/routes"
)

func main() {
	// Load configuration
	cfg := config.Load()
	log.Printf("Starting Bill Tracker API in %s mode", cfg.Env)

	// Connect to database
	db := database.Connect(cfg)

	// Run migrations
	if err := database.Migrate(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Setup router
	router := routes.SetupRouter(db)

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
