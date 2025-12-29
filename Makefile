.PHONY: help build run test clean fmt lint install-deps docker-up docker-down migrate

# Variables
APP_NAME=go-training-api
MAIN_PATH=cmd/api/main.go
BINARY_NAME=bin/$(APP_NAME)
GO=go
DOCKER_COMPOSE=docker-compose
# --- Database URL Construction from .env ---
# We extract values using grep/cut to avoid hardcoding
DB_USER=$(shell grep DB_USER .env | cut -d '=' -f2)
DB_PASSWORD=$(shell grep DB_PASSWORD .env | cut -d '=' -f2)
DB_HOST=$(shell grep DB_HOST .env | cut -d '=' -f2)
DB_PORT=$(shell grep DB_PORT .env | cut -d '=' -f2)
DB_NAME=$(shell grep DB_NAME .env | cut -d '=' -f2)
DB_SSLMODE=$(shell grep DB_SSLMODE .env | cut -d '=' -f2)

DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)
MIGRATIONS_PATH=migrations

# ... (keep your help, build, and run targets) ...

# --- Migrations ---

migrate-up:
	@echo "Running migrations up..."
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

migrate-down:
	@echo "Rolling back 1 migration..."
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down 1

migrate-force:
	@read -p "Enter version to force: " version; \
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" force $$version

migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $$name
	
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