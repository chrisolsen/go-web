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

Podman is required

### Go dev tools

```bash
go install golang.org/x/tools/gopls@latest
go install github.com/nametake/golangci-lint-langserver@latest
go install github.com/go-delve/delve/cmd/dlv@latest
```

### Tailwind
Either install tailwind as a global npm lib via `npm i -g tailwindcss` or download the binary from
the [releases page](https://github.com/tailwindlabs/tailwindcss/releases) and save it in the root folder

### Tools

These are only required if you are not using the Podman containers

```bash
# Dev only required libs 
go install go.uber.org/mock/mockgen@latest
go install github.com/air-verse/air@latest
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
go install -tags 'postgresql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### Libs

```bash
# ensure all go.mod libs are installed
go get ./...
```

### Start dev server
```bash

# shell 1
make dev-compose
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
