package main

import (
	"fmt"
	"log"
)

func main() {
	// Init configuration
	config, err := LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	fmt.Printf("Service will run on port: %s\n", config.Server.Port)

	// Init logger
	if err := ConfigureLogger(config.Logging.Level); err != nil {
		log.Fatalf("Failed to configure logger: %v", err)
	}

	log.Println("Logger initialized")

	// Example usage of logging at different levels
	Info("Application started")
	Debug("This is a debug message")
	Error("An error occurred")
}
