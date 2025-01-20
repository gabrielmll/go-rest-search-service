package logger

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

// TestConfigureLogger_ValidLevels ensures that valid log levels are correctly configured.
func TestConfigureLogger_ValidLevels(t *testing.T) {
	tests := []struct {
		name        string
		level       string
		expected    LogLevel
		shouldError bool
	}{
		{"debug", "debug", DebugLevel, false},
		{"info", "info", InfoLevel, false},
		{"error", "error", ErrorLevel, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := ConfigureLogger(tt.level)

			// Check if error is expected
			if (err != nil) != tt.shouldError {
				t.Fatalf("expected error: %v, got: %v", tt.shouldError, err)
			}

			// Ensure log level is correctly set
			if CurrentLogLevel() != tt.expected {
				t.Errorf("expected log level: %v, got: %v", tt.expected, CurrentLogLevel())
			}
		})
	}
}

// TestConfigureLogger_InvalidLevel verifies that providing an invalid log level
// returns an appropriate error and does not set the log level.
func TestConfigureLogger_InvalidLevel(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	err := ConfigureLogger("invalid")
	if err == nil {
		t.Fatal("expected an error but got none")
	}

	expectedError := "invalid log level: invalid"
	if err.Error() != expectedError {
		t.Errorf("expected error %q, got %q", expectedError, err.Error())
	}
}

// TestLogMessage_Debug checks that a debug-level message is correctly logged
// when the logger is set to the debug level.
func TestLogMessage_Debug(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	_ = ConfigureLogger("debug") // Set log level to debug

	Debug("This is a debug message")
	output := buf.String()

	if !strings.Contains(output, "[DEBUG] This is a debug message") {
		t.Errorf("expected debug log message, got %q", output)
	}
}

// TestLogMessage_Info ensures that an info-level message is correctly logged
// when the logger is set to the info level.
func TestLogMessage_Info(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	_ = ConfigureLogger("info") // Set log level to info

	Info("This is an info message")
	output := buf.String()

	if !strings.Contains(output, "[INFO] This is an info message") {
		t.Errorf("expected info log message, got %q", output)
	}
}

// TestLogMessage_Error ensures that an error-level message is correctly logged
// regardless of the configured log level.
func TestLogMessage_Error(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	_ = ConfigureLogger("error") // Set log level to error

	Error("This is an error message")
	output := buf.String()

	if !strings.Contains(output, "[ERROR] This is an error message") {
		t.Errorf("expected error log message, got %q", output)
	}
}

// TestLogMessage_LevelFiltering verifies that messages below the configured log level
// are not logged, while messages at or above the configured level are logged.
func TestLogMessage_LevelFiltering(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	_ = ConfigureLogger("info") // Set log level to info

	Debug("This debug message should not appear")
	Info("This info message should appear")
	Error("This error message should appear")

	output := buf.String()

	if strings.Contains(output, "This debug message should not appear") {
		t.Error("debug message logged at info level")
	}

	if !strings.Contains(output, "This info message should appear") {
		t.Error("info message not logged at info level")
	}

	if !strings.Contains(output, "This error message should appear") {
		t.Error("error message not logged at info level")
	}
}
