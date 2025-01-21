# Makefile for go-rest-search-service project

# Set the Go binary and test flags
GO := go
TEST_FLAGS := -v

# Target to run tests
.PHONY: test
test:
	$(GO) test $(TEST_FLAGS) ./...

# Target to run tests with coverage
.PHONY: test-coverage
test-coverage:
	$(GO) test $(TEST_FLAGS) -cover ./...

# Target to run tests and generate a test report in a file
.PHONY: test-report
test-report:
	$(GO) test $(TEST_FLAGS) -json ./... > test_report.json

# Target to format Go files
.PHONY: fmt
fmt:
	$(GO) fmt ./...

# Target to run the project
.PHONY: run
run:
	$(GO) run ./cmd

# Target to build the project
.PHONY: build
build:
	$(GO) build -o myapp

# Clean the build files
.PHONY: clean
clean:
	rm -f myapp
