ifneq (,$(wildcard .env))
include .env
export DATABASE_URL
endif

DATABASE_URL ?= postgres://app:local_dev_password@localhost:5432/students_db?sslmode=disable

postgres:
	docker run --name student-postgres \
		-p 5432:5432 \
		-e POSTGRES_USER=app \
		-e POSTGRES_PASSWORD=local_dev_password \
		-d postgres:18.4-alpine

createdb:
	docker exec -it student-postgres createdb --username=app --owner=app students_db

dropdb:
	docker exec -it student-postgres dropdb students_db

migrateup:
	migrate -path db/migration -database "$(DATABASE_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DATABASE_URL)" -verbose down

migrate-up: migrateup

migrate-down: migratedown

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run cmd/api/main.go

docker-build:
	docker build -t student-api:latest .

docker-up:
	docker compose -f docker-compose.prod.yml up -d --build

docker-down:
	docker compose -f docker-compose.prod.yml down

.PHONY: postgres createdb dropdb migrateup migratedown migrate-up migrate-down sqlc test server docker-build docker-up docker-down