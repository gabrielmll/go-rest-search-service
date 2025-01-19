package tests

import (
	"go-rest-search-service/internal/config"
	"os"
	"strings"
	"testing"
)

// TestLoadConfig_ValidFile validates if a config.yaml is properly loaded.
func TestLoadConfig_ValidFile(t *testing.T) {
	// Create a temporary valid config file
	tempFile, err := os.CreateTemp("", "valid-config-*.yaml")
	if err != nil {
		t.Fatalf("failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	validConfigContent := `
server:
  port: "8080"
logging:
  level: "info"
file:
  path: "input.txt"
`
	if _, err := tempFile.WriteString(validConfigContent); err != nil {
		t.Fatalf("failed to write to temporary file: %v", err)
	}

	// Load the config file
	cfg, err := config.LoadConfig(tempFile.Name())
	if err != nil {
		t.Fatalf("unexpected error while loading config: %v", err)
	}

	// Validate the config values
	if cfg.Server.Port != "8080" {
		t.Errorf("expected server port '8080', got '%s'", cfg.Server.Port)
	}
	if cfg.Logging.Level != "info" {
		t.Errorf("expected logging level 'info', got '%s'", cfg.Logging.Level)
	}
	if cfg.File.Path != "input.txt" {
		t.Errorf("expected file path 'input.txt', got '%s'", cfg.File.Path)
	}
}

// TestLoadConfig_FileNotFound validates if an invalid file path will raise the expected error.
func TestLoadConfig_FileNotFound(t *testing.T) {
	_, err := config.LoadConfig("non-existent-file.yaml")
	if err == nil {
		t.Fatal("expected an error for a non-existent file, but got none")
	}

	expectedError := "failed to open config file"
	if err != nil && !contains(err.Error(), expectedError) {
		t.Errorf("expected error containing '%s', got '%v'", expectedError, err)
	}
}

// TestLoadConfig_EmptyFile validates if an empty config YAML file will raise the expected error.
func TestLoadConfig_EmptyFile(t *testing.T) {
	// Create an empty temporary file
	tempFile, err := os.CreateTemp("", "empty-config-*.yaml")
	if err != nil {
		t.Fatalf("failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	_, err = config.LoadConfig(tempFile.Name())
	if err == nil {
		t.Fatal("expected an error for an empty config file, but got none")
	}

	expectedError := "config file is empty or improperly formatted"
	if err != nil && err.Error() != expectedError {
		t.Errorf("expected error '%s', got '%v'", expectedError, err)
	}
}

// TestLoadConfig_InvalidFormat validates if an invalid formated config YAML file will raise the expected error.
func TestLoadConfig_InvalidFormat(t *testing.T) {
	// Create a temporary file with invalid YAML content
	tempFile, err := os.CreateTemp("", "invalid-config-*.yaml")
	if err != nil {
		t.Fatalf("failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// To spot the invalid format, notice the missing quotes after "info
	invalidConfigContent := `
server:
  port: "8080"
logging:
  level: "info
`
	if _, err := tempFile.WriteString(invalidConfigContent); err != nil {
		t.Fatalf("failed to write to temporary file: %v", err)
	}

	_, err = config.LoadConfig(tempFile.Name())
	if err == nil {
		t.Fatal("expected an error for invalid YAML content, but got none")
	}

	expectedError := "failed to decode config file"
	if err != nil && !contains(err.Error(), expectedError) {
		t.Errorf("expected error containing '%s', got '%v'", expectedError, err)
	}
}

// Helper function to check if a string contains another string
func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr || len(s) > len(substr) && strings.Contains(s, substr)
}
