package main

import (
	"log"

	"github.com/moshfiq123456/ums-be/internal/app"
	"github.com/moshfiq123456/ums-be/internal/config"
	"github.com/moshfiq123456/ums-be/internal/database"
	"github.com/moshfiq123456/ums-be/internal/logger"
)

func main() {
	cfg := config.LoadConfig()
	
	// Initialize file-based logger
	logger.InitLogger(cfg.LogFilePath)
	
	// Run database migrations
	log.Println("Running database migrations...")
	migrationService := database.NewMigrationService(cfg.GetDatabaseURL())
	if err := migrationService.RunMigrations(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	
	// Start the server
	server := app.NewServer(cfg)
	server.Start()
}