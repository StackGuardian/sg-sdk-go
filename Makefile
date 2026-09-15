# Makefile for StackGuardian Go SDK

.PHONY: all build test fmt vet lint clean

# Default target
all: fmt vet build

# Format the Go SDK code
fmt:
	gofmt -w .
	goimports -w .

# Run go vet
vet:
	go vet ./...

# Build the SDK (verify compilation)
build:
	go build ./...

# Run tests
test:
	go test -v -race -cover ./...

# Run linter (requires golangci-lint)
lint:
	golangci-lint run

# Clean build artifacts
clean:
	go clean ./...
	rm -rf dist/

# Install development dependencies
install-dev-tools:
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run examples (requires SG_API_TOKEN environment variable)
run-example-basic:
	go run examples/basic/main.go

run-example-workflow:
	go run examples/workflow_run/main.go
