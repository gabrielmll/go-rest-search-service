package main

import (
	"log"
	"net/http"

	"go-rest-search-service/internal/api"
	"go-rest-search-service/internal/config"
	"go-rest-search-service/internal/logger"
	"go-rest-search-service/internal/utils"
)

// main initializes the configuration, logger, data loading, and starts the HTTP server.
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
	logger.Info("Logger initialized. Log level set to %v", config.Logging.Level)

	// Load numbers from the file to a slice
	numbers, err := utils.LoadNumbersFile(config.File.Path)
	if err != nil {
		log.Fatalf("Failed to load numbers from file: %v", err)
	}
	logger.Info("Numbers loaded from %s\n", config.File.Path)

	// Init the middleware
	handler := http.HandlerFunc(api.EndpointHandler(numbers))

	// Register api endpoint
	http.Handle("/endpoint/", api.LoggerMiddleware(handler))

	logger.Info("Starting server on port: %s\n", config.Server.Port)
	log.Fatal(http.ListenAndServe(":"+config.Server.Port, nil))
}
