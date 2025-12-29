.PHONY: help build run test clean fmt lint install-deps docker-up docker-down migrate

# Variables
APP_NAME=go-training-api
MAIN_PATH=cmd/api/main.go
BINARY_NAME=bin/$(APP_NAME)
GO=go
DOCKER_COMPOSE=docker-compose

help:
	@echo "Available commands:"
	@echo "  make install-deps    - Install dependencies"
	@echo "  make build          - Build the application"
	@echo "  make run            - Run the application"
	@echo "  make dev            - Run in development mode with hot reload"
	@echo "  make test           - Run tests"
	@echo "  make fmt            - Format code"
	@echo "  make lint           - Run linter"
	@echo "  make clean          - Clean build artifacts"
	@echo "  make docker-up      - Start PostgreSQL container"
	@echo "  make docker-down    - Stop PostgreSQL container"

install-deps:
	$(GO) mod download
	$(GO) mod tidy

build: install-deps
	$(GO) build -o $(BINARY_NAME) $(MAIN_PATH)
	@echo "Build complete: $(BINARY_NAME)"

run: build
	./$(BINARY_NAME)

dev:
	@command -v air >/dev/null 2>&1 || (echo "Installing air..." && $(GO) install github.com/cosmtrek/air@latest)
	air

test:
	$(GO) test -v -cover ./...

fmt:
	$(GO) fmt ./...

lint:
	@command -v golangci-lint >/dev/null 2>&1 || (echo "Installing golangci-lint..." && $(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	golangci-lint run ./...

clean:
	rm -f $(BINARY_NAME)
	$(GO) clean

docker-up:
	$(DOCKER_COMPOSE) up -d postgres

docker-down:
	$(DOCKER_COMPOSE) down

migrate:
	@echo "Run migrations using your preferred migration tool"