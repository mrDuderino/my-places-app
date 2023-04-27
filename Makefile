API_BINARY=myPlacesApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_api
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_broker: builds the broker binary as a linux executable
build_api:
	@echo "Building api binary..."
	env GOOS=linux CGO_ENABLED=0 go build -o build/${API_BINARY} ./cmd/my-places-api
	@echo "Done!"

.PHONY: migrate-create migrate-up migrate-down connect-db

migrate-create:
	migrate create -ext sql -dir ./schema -seq init

migrate-up:
	migrate -path ./schema -database 'postgres://postgres:qwertysha@localhost:5432/postgres?sslmode=disable' -verbose up

migrate-down:
	migrate -path ./schema -database 'postgres://postgres:qwertysha@localhost:5432/postgres?sslmode=disable' -verbose down

connect:
	docker exec -it my-places-db /bin/bash