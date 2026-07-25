# Common
APP_NAME = chrisolsen-goweb
MIGRATION_PATH = ./db/migration

# Init
init:
	./scripts/init.sh

# Dev

dev-setup:
	go install golang.org/x/tools/gopls@latest
	go install github.com/nametake/golangci-lint-langserver@latest
	go install github.com/go-delve/delve/cmd/dlv@latest
	go get ./...

dev-compose:
	podman compose -f .docker/docker-compose.dev.yml up

dev-css:
	tailwindcss -i app/styles/input.css -o app/static/output.css --watch

dev-serve:
	podman exec docker_app_1 air -c .air.toml

# Build

build:
	go build -o ./bin/public/app ./app


build-prod:
	podman compose up -f .docker/Dockerfile.prod


# Testing


test:
	go test $(APP_NAME)


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
