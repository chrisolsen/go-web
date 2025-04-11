# README

## Prerequisits
- Docker
- DockerCompose


## Setup

> ** Perform a search and replace of `chrisolsen-goweb` with the name of your app.

### Tools (non-devontainer only)
```bash
# Run on local machine
go install go.uber.org/mock/mockgen@latest
go install github.com/air-verse/air@latest
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Development

### Start dev server
```bash
go mod install  # non-devcontainer only
air
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
