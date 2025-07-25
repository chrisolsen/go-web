# Common
APP_NAME = chrisolsen-goweb 
MIGRATION_PATH = ./db/migration

dev-containers-up:
	docker-compose -f .docker/docker-compose.dev.yml up

dev: 
	docker exec -it docker-app-1 air

# Build

public-build:
	go build -o ./bin/app ./apps/public/main.go


public-build-prod:
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
