package main

import (
	"github.com/moshfiq123456/ums-be/internal/app"
	"github.com/moshfiq123456/ums-be/internal/config"
	"github.com/moshfiq123456/ums-be/internal/logger"
)

func main() {
	cfg := config.LoadConfig()
	
	// Initialize file-based logger
	logger.InitLogger(cfg.LogFilePath)
	
	server := app.NewServer(cfg)
	server.Start()
}