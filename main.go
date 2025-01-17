package main

import (
	"log"
)

func main() {
	// Init configuration
	config, err := LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Init logger
	if err := ConfigureLogger(config.Logging.Level); err != nil {
		log.Fatalf("Failed to configure logger: %v", err)
	}

	Info("Logger initialized")
	Info("Service will run on port: %s\n", config.Server.Port)

	// Example usage of logging at different levels
	Info("Application started")
	Debug("This is a debug message")
	Error("An error occurred")
}
