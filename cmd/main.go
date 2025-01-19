package main

import (
	"log"
	"net/http"

	"go-rest-search-service/internal/config"
	"go-rest-search-service/internal/logger"
	"go-rest-search-service/pkg/api"
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

	http.HandleFunc("/test/", api.EndpointHandler())

	logger.Info("Starting server on port: %s\n", config.Server.Port)
	log.Fatal(http.ListenAndServe(":"+config.Server.Port, nil))
}
