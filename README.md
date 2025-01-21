# Go REST Search Service

This project implements a simple Go-based REST API that performs binary search on a sorted list of integers. It allows users to search for a specific integer, and if it's not found, it returns the closest integer within a 10% margin of the target.

## Features

- **Binary Search API**: Efficient search for exact or closest values within a sorted list.
- **Error Handling**: Provides meaningful error messages for invalid input or missing values.
- **Test Automation**: Fully tested with unit tests.
- **Configuration File**: Uses a YAML-based configuration file for easy environment setup.
- **Logging**: Configurable logging to monitor application behavior.

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/gabrielmll/go-rest-search-service.git
   cd go-rest-search-service

2. **Install dependencies**:
The project uses Go modules for dependency management, so simply run:
    ```bash
    go mod tidy

3. **Configuration**:
- Make sure to create a config.yaml file based on the provided sample (config.sample.yaml).
- Set the appropriate file path for the numbers file and server configurations.


## Running the Service
After configuring the application, run the service with:
    ```bash
    make run

If logger is set to show INFO, you should see this output: 
    ```console
    2025-01-21 20:45:35 [INFO] Logger initialized. Log level set to debug
    2025-01-21 20:45:36 [INFO] Numbers loaded from input.txt
    2025-01-21 20:45:36 [INFO] Starting server on port: 8080





## File Structure
```
/go-rest-search-service
├── cmd/
│   └── main.go            # Entry point for the API server
├── internal/
│   ├── api/               # API handlers and routing
│   ├── config/            # Configuration loading
|   ├── logger/            # Logging utilities
│   ├── service/           # Business logic (binary search)
│   └── utils/             # Helper functions (file reading)
├── config.yaml            # Configuration file
├── Makefile               # Build and test automation
├── README.md              # Project documentation
└── go.mod                 # Go module file