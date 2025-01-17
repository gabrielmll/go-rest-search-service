package main

import (
	"log"

	"go-rest-search-service/internal/config"
	"go-rest-search-service/internal/logger"
)

func main() {
	// Init configuration
	config, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Init logger
	if err := logger.ConfigureLogger(config.Logging.Level); err != nil {
		log.Fatalf("Failed to configure logger: %v", err)
	}

	logger.Info("Logger initialized")
	logger.Info("Service will run on port: %s\n", config.Server.Port)

	// Example usage of logging at different levels
	logger.Info("Application started")
	logger.Debug("This is a debug message")
	logger.Error("An error occurred")
}
