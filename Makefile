# Common
APP_NAME = chrisolsen-goweb 
MIGRATION_PATH = ./db/migration


# Build

build:
	go build -o ./bin/web ./cmd/server/main.go


build-prod:
	docker-compose up -f Dockerfile.prod
.PHONY: build-prod


# Testing


test:
	go test \
		$(APP_NAME)/internal/transport/http \
		$(APP_NAME)/cmd/server
.PHONY: test


benchmark:
	go benchmark $(APP_NAME)/benchmarks
.PHONY: test


# Database

db-create:
	docker exec -it db createdb --username=postgres --owner=postgres $(APP_NAME)
.PHONY: db-create


db-drop:
	docker exec -it db dropdb --username=postgres --owner=postgres $(APP_NAME)
.PHONY: db-drop


# Migrations


# make migrate-create name=[migration_name]
migrate-create:
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq $(name)
.PHONY: migrate-create


migrate:
	migrate -path $(MIGRATION_PATH) -database "postgresql://postgres:postgres@localhost:5432/$(APP_NAME)?sslmode=disable" -verbose up
.PHONY: migrate


# Sqlc - Generate the Go files for the migrations


sqlc:
	sqlc generate
.PHONY: sqlc
