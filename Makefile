.PHONY: build run test clean up down restart

# Build the Go application
build:
	go build -o bin/app ./cmd

# Run the application
run:
	go run ./cmd

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Start services with Docker Compose
up:
	docker compose -f docker/docker-compose.local.yaml up -d

# Stop services with Docker Compose
down:
	docker compose -f docker/docker-compose.local.yaml down

# Restart services with Docker Compose
restart:
	docker compose -f docker/docker-compose.local.yaml down
	docker compose -f docker/docker-compose.local.yaml up -d