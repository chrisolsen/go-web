# Common
APP_NAME = chrisolsen-goweb 
MIGRATION_PATH = ./db/migration

docker-up:
	docker-compose -f .docker/docker-compose.dev.yml up

run:
	air

# Build

build:
	go build -o ./bin/web ./cmd/server/main.go


build-prod:
	docker-compose up -f .docker/Dockerfile.prod


# Testing


test:
	go test \
		$(APP_NAME)/internal/transport/http \
		$(APP_NAME)/cmd/server


benchmark:
	go benchmark $(APP_NAME)/benchmarks


# Database

db-create:
	docker exec -it db createdb --username=postgres --owner=postgres $(APP_NAME)


db-drop:
	docker exec -it db dropdb --username=postgres --owner=postgres $(APP_NAME)


# Migrations


# make migrate-create name=[migration_name]
migrate-create:
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq $(name)


migrate:
	migrate -path $(MIGRATION_PATH) -database "postgresql://postgres:postgres@localhost:5432/$(APP_NAME)?sslmode=disable" -verbose up


# Sqlc - Generate the Go files for the migrations


sqlc:
	sqlc generate
