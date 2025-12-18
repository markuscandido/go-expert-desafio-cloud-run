.PHONY: run build test clean docker-build docker-run

# Load .env file if exists
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Application
run:
	go run cmd/api/main.go

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin/server ./cmd/api

test:
	go test ./... -v

test-cover:
	go test ./... -cover -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

clean:
	rm -rf bin/ coverage.out coverage.html

# Docker
docker-build:
	docker build -t weather-api:latest .

docker-run:
	docker run -p 8080:8080 -e WEATHER_API_KEY=$(WEATHER_API_KEY) weather-api:latest

# Docker Compose
up:
	docker-compose up --build

down:
	docker-compose down

# Development
tidy:
	go mod tidy

fmt:
	go fmt ./...

lint:
	golangci-lint run ./...
