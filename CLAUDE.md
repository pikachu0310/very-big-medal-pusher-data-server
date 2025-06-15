# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a VRChat Medal Pusher Game Data API Server that collects gameplay statistics via HMAC-signed requests and provides ranking APIs. The server uses OpenAPI-first development with oapi-codegen for type-safe API generation.

## Development Commands

### Start Development Environment
```bash
docker compose watch  # Auto-rebuilds on file changes
```
- API: http://localhost:8080/
- DB Admin: http://localhost:8081/
- MariaDB: localhost:3306

### Code Generation (OpenAPI)
```bash
make oapi              # Generate both server and models
make generate-server   # Generate server interfaces only
make generate-models   # Generate model structs only
```

### Testing and Quality
```bash
make test             # Run all tests
make test-unit        # Unit tests only
make test-integration # Integration tests only
make lint             # Run golangci-lint with --fix
```

### Build and Dependencies
```bash
make build    # Compile binary
make mod      # Download dependencies
```

## Architecture

### OpenAPI-First Development
1. Edit `openapi/openapi.yaml` to define API changes
2. Run `make oapi` to generate Go code
3. Implement `ServerInterface` methods in `internal/handler/`
4. Add database migrations in `internal/migration/` if needed

### Project Structure
- `internal/handler/`: HTTP handlers for v1, v2, v3 APIs + admin
- `internal/repository/`: Database access layer using sqlx
- `internal/domain/`: Data transformation between API and DB models
- `internal/migration/`: Sequential SQL migrations using Goose
- `internal/pkg/config/`: Environment-based configuration
- `openapi/`: Generated server interfaces and models
- `tools/`: Code generation configuration
- `web/admin/`: Bootstrap admin dashboard

### API Versioning
- v1: Original GET-based API
- v2: Enhanced with POST endpoints
- v3: Latest version with extended functionality
- All versions maintained for backward compatibility

### Database Migrations
Uses Goose with embedded SQL files:
- Sequential numbering: `1_schema.sql`, `2_add_column.sql`, etc.
- Auto-migration on startup in `main.go`
- UTF8MB4 collation for MariaDB/MySQL

### Security
- HMAC-SHA256 signature verification for data submissions
- Multiple secret keys for different operations (SAVE, LOAD, etc.)
- JWT authentication for admin interface

## Adding New Parameters

Follow `HowToAddNewParmsToServer.md`:
1. Add parameter to `openapi/openapi.yaml`
2. Create migration SQL in `internal/migration/`
3. Run `make oapi` to generate code
4. Update `internal/domain/data.go`
5. Update `internal/repository/data.go` InsertGameData method

## Environment Variables

Required for full functionality:
```bash
# Admin authentication
ADMIN_USER_ID=admin
ADMIN_PASSWORD_HASH=hashed_password
ADMIN_JWT_SECRET=jwt_secret

# Game data security  
SECRET_KEY=main_secret
SAVE=save_secret
LOAD=load_secret

# Database (auto-detected for local vs NeoShowcase)
DB_HOST=localhost
DB_USER=root
DB_PASSWORD=pass
DB_NAME=app
```

## Technologies

- **Framework**: Echo v4 (HTTP), sqlx (DB), Goose (migrations)
- **Code Gen**: oapi-codegen for OpenAPI
- **Security**: golang-jwt, HMAC-SHA256
- **Caching**: motoki317/sc with TTL
- **Development**: Docker Compose with watch mode
- **Linting**: golangci-lint