package logger

import (
	"fmt"
	"io"
	"log"
	"strings"
	"time"
)

// LogLevel type defines the available log levels
type LogLevel int

const (
	// Log levels
	DebugLevel LogLevel = iota
	InfoLevel
	ErrorLevel
)

// Global variable to store the current log level
var currentLogLevel LogLevel

// ConfigureLogger sets the log level based on the config
func ConfigureLogger(level string) error {
	// Set the log level based on the input string
	switch strings.ToLower(level) {
	case "debug":
		currentLogLevel = DebugLevel
	case "info":
		currentLogLevel = InfoLevel
	case "error":
		currentLogLevel = ErrorLevel
	default:
		return fmt.Errorf("invalid log level: %s", level)
	}

	log.SetFlags(0)
	log.SetPrefix("")

	return nil
}

// logMessage logs a message with my custom format
func logMessage(level string, format string, v ...interface{}) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	message := fmt.Sprintf(format, v...)
	log.Printf("%s [%s] %s", currentTime, level, message)
}

// Debug logs a message with the Debug level
func Debug(format string, v ...interface{}) {
	if currentLogLevel <= DebugLevel {
		logMessage("DEBUG", format, v...)
	}
}

// Info logs a message with the Info level
func Info(format string, v ...interface{}) {
	if currentLogLevel <= InfoLevel {
		logMessage("INFO", format, v...)
	}
}

// Error logs a message with the Error level
func Error(format string, v ...interface{}) {
	if currentLogLevel <= ErrorLevel {
		logMessage("ERROR", format, v...)
	}
}

// CurrentLogLevel returns the current log level
func CurrentLogLevel() LogLevel {
	return currentLogLevel
}

// SetOutput allows redirecting log output for testing
func SetOutput(w io.Writer) {
	log.SetOutput(w)
}
