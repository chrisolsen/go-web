# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Setup and Installation
```bash
./init.sh                           # Initial project setup
go mod tidy                         # Install Go dependencies (non-devcontainer only)
```

### Development Server
```bash
air                                 # Start development server with hot reload
make dev                           # Start dev server in Docker container
make dev-containers-up             # Bring up development containers
```

### Build Commands
```bash
make public-build                  # Build public app to ./bin/app
make public-build-prod             # Build production Docker image
```

### Testing
```bash
make test                          # Run tests for admin and public apps
make benchmark                     # Run benchmarks
go test ./...                      # Run all tests in project
```

### Database Operations
```bash
make db-create                     # Create SQLite database
make migrate-create name=[name]    # Create new migration with given name
make migrate                       # Run all pending migrations
make sqlc                          # Generate Go code from SQL queries
```

## Architecture Overview

This is a Go web application with a modular architecture supporting multiple apps (public and admin) that share common infrastructure.

### Project Structure
- **apps/**: Contains the main applications
  - `apps/base.go`: Shared base application with HTTP server setup and graceful shutdown
  - `apps/public/`: Public-facing web application
  - `apps/admin/`: Admin interface application
- **internal/**: Private application code
  - `internal/db/`: Database layer with SQLC-generated code
  - `internal/handlers/`: HTTP request handlers organized by feature
  - `internal/services/`: Business logic services (auth, email, payment, logging, health)
- **db/**: Database files
  - `db/migration/`: SQL migration files
  - `db/query/`: SQL query definitions for SQLC

### Key Architectural Patterns
- **Service Layer**: Business logic is encapsulated in services with interface-based design
- **SQLC Integration**: Database queries are written in SQL and Go code is generated automatically
- **Multiple Apps**: Shared base functionality with app-specific implementations
- **Graceful Shutdown**: HTTP servers implement proper graceful shutdown handling
- **SQLite Database**: Uses SQLite with migration support via golang-migrate

### Service Interfaces
All services follow interface-based design patterns:
- `Authenticator`: Authentication service
- `Emailer`: Email service  
- `Paymenter`: Payment processing service
- `Logger`: Logging service

### Database Workflow
1. Create migrations with `make migrate-create name=[name]`
2. Add SQL to migration files in `db/migration/`
3. Run migrations with `make migrate`
4. Add queries to `db/query/` SQL files
5. Generate Go code with `make sqlc`

### Development Tools Required
- Docker and Docker Compose
- Go 1.24+
- air (for hot reload)
- sqlc (for code generation)
- golang-migrate (for database migrations)
- mockgen (for testing)