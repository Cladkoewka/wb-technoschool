.PHONY: dev docker-up docker-down migrate-up migrate-down export-env

DB_DSN = "postgres://level0user:level0password@localhost:5435/level0db?sslmode=disable"

dev:
	go run cmd/main.go

docker-up:
	docker compose up -d

docker-down:
	docker compose down

migrate-up:
	goose -dir migrations/postgres postgres $(DB_DSN) up

migrate-down:
	goose -dir migrations/postgres postgres $(DB_DSN) down

start: docker-up migrate-up dev