# Common
APP_NAME = chrisolsen-goweb 
MIGRATION_PATH = ./db/migration

# Init
init:
	./scripts/init.sh

# Dev

dev:
	docker-compose -f .docker/docker-compose.dev.yml up

public: 
	docker exec docker-app-1 air -c .air.public.toml

admin: 
	docker exec docker-app-1 air -c .air.admin.toml

# Build

build-dev:
	go build -o ./bin/public ./apps/public/main.go
	go build -o ./bin/admin ./apps/admin/main.go


build-prod:
	docker-compose up -f .docker/Dockerfile.prod


# Testing


test:
	go test \
		$(APP_NAME)/apps/admin \
		$(APP_NAME)/apps/public


benchmark:
	go benchmark $(APP_NAME)/benchmarks


# Database

db-create:
	sqlite3 database.db


# Migrations


# make migrate-create name=[migration_name]
migrate-create:
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq $(name)


migrate:
	migrate -path $(MIGRATION_PATH) -database "./database.db" -verbose up


# Sqlc - Generate the Go files for the migrations


sqlc:
	sqlc generate
