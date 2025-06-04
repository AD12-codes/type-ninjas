# Makefile for Go project
.PHONY: all build run clean deps migration-create migration-up migration-down sqlc-generate

# load .env
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Variables
APP_NAME := type-ninjas
BUILD_DIR := bin
MIGRATIONS_PATH := file://./db/migrations


# Default target
all: build

# Build the application
build:
	@echo "Building the application..."
	@go build -o $(BUILD_DIR)/main cmd/main.go

# Run the application
dev:
	@echo "Running the application..."
	@air

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)

deps:
	go mod tidy

# migration commands
migration-create:
ifndef name
	$(error name is required. Usage: make migration-create name=create_users_table)
endif
	migrate create -ext sql -dir ./db/migrations -seq $(name)

migration-up:
	migrate -source $(MIGRATIONS_PATH) -database $(POSTGRESQL_URL) up

migration-down:
	migrate -source $(MIGRATIONS_PATH) -database $(POSTGRESQL_URL) down

# sqlc commands
sqlc-generate:
	sqlc generate

# tests
test:
	go test ./...
