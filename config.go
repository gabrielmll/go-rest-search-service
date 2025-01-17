package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config struct holds the configuration values
type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Logging struct {
		Level string `yaml:"level"`
	} `yaml:"logging"`
	File struct {
		Path string `yaml:"path"`
	} `yaml:"file"`
}

// LoadConfig loads the configuration from the provided YAML file
func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		if err.Error() == "EOF" {
			return nil, fmt.Errorf("config file is empty or improperly formatted")
		}
		return nil, fmt.Errorf("failed to decode config file: %v", err)
	}

	return &config, nil
}
