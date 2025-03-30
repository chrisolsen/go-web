# Common
PKG_NAME = __PKG__ 
DB_NAME = __DB__
MIGRATION_PATH = ./db/migration

dev: docker-dev
	echo Running developement server...
.PHONY: dev

dev-local:
	modd
.PHONY: dev-local

run:
	go run cmd/server/main.go
.PHONY: run


# Docker


docker-dev:
	docker-compose -f "docker-compose.yml" up -d
.PHONY: docker-dev


docker-build:
	docker-compose up -f Dockerfile.prod
.PHONY: docker-build


# Testing


test:
	go test \
		$(PKG_NAME)/internal/transport/http \
		$(PKG_NAME)/cmd/server
.PHONY: test


benchmark:
	go benchmark $(PKG_NAME)/benchmarks
.PHONY: test


# Database

db-create:
	docker exec -it db createdb --username=postgres --owner=postgres $(DB_NAME)
.PHONY: db-create


db-drop:
	docker exec -it db dropdb --username=postgres --owner=postgres $(DB_NAME)
.PHONY: db-drop


# Migrations


# make migrate-create name=[migration_name]
migrate-create:
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq $(name)
.PHONY: migrate-create


migrate:
	migrate -path $(MIGRATION_PATH) -database "postgresql://postgres:postgres@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose up
.PHONY: migrate


# Sqlc - Generate the Go files for the migrations


sqlc:
	sqlc generate
.PHONY: sqlc
