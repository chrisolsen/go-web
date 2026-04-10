# README

## Prerequisits

- Docker
- DockerCompose

## Initialize new project 

```bash
# updates the go module name 
./init.sh 
```

## Development

A docker development environment exists that contains the various libs (see below) required on the development side.

### Tools 

```bash
# Dev only required libs 
go install go.uber.org/mock/mockgen@latest
go install github.com/air-verse/air@latest
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
go install -tags 'postgresql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# ensure all go.mod libs are installed
go get ./...
```

### Start dev server
```bash

# shell 1
make dev-up
# shell 2
make dev
```

### Common commands
```bash
make db-create			        # create the database
make migrate-create name=[name]	# create a `name` migration, add sql code to migration files
make migrate			        # execute migration
make sqlc			            # creates go code
```

## Tips
- https://www.youtube.com/watch?v=wxkEQxvxs3w
