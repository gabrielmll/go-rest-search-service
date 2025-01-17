package main

import (
	"fmt"
	"log"
	"strings"
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

	// Set the log output format
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// Output a message indicating the configured log level
	log.Printf("Log level set to %s", level)
	return nil
}

// LogDebug logs a message with the Debug level
func Debug(v ...interface{}) {
	if currentLogLevel <= DebugLevel {
		log.SetPrefix("[DEBUG] ")
		log.Println(v...)
	}
}

// LogInfo logs a message with the Info level
func Info(v ...interface{}) {
	if currentLogLevel <= InfoLevel {
		log.SetPrefix("[INFO] ")
		log.Println(v...)
	}
}

// LogError logs a message with the Error level
func Error(v ...interface{}) {
	if currentLogLevel <= ErrorLevel {
		log.SetPrefix("[ERROR] ")
		log.Println(v...)
	}
}
