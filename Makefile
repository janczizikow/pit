include .env
export

.PHONY: build migration migrate rollback clean

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## run: runs the cmd/api application
.PHONY: run
run:
	go run ./cmd/api

## build
.PHONY: build
build:
	@echo "building cmd/api..."
	go build -o=./bin/api ./cmd/api

## migration name=$1: creates a new database migration
.PHONY: migration
migration:
	migrate create -ext sql -dir ./internal/database/migrations $(name)

## migrate: applies all up database migrations
.PHONY: migrate
migrate:
	migrate -path ./internal/database/migrations -database=$(DB_DSN) up

## rollback
.PHONY: rollback
rollback:
	migrate -path ./internal/database/migrations -database=$(DB_DSN) down

## psql: connects to the database using psql
.PHONY: psql
psql:
	psql $(DB_DSN)

## test: run unit tests
test:
	go test -race -cover ./...

## clean
.PHONY: clean
clean:
	@echo "cleaning cmd/api"
	@rm ./bin/api
