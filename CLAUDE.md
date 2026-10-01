# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**Hotel Reservation API** — A REST API for managing hotel reservations, currently implementing user management. Built with Go, Fiber (HTTP framework), and MongoDB.

## Architecture

The codebase follows a **layered architecture**:

- **`main.go`** — Server setup: initializes MongoDB client, wires up handlers, and starts the Fiber HTTP server on port 5000
- **`types/`** — Domain models and request/response types
  - `User` struct with validation: email, names (2-50 chars), age, password (hashed with bcrypt at cost 12)
  - `CreateUserParam` — request body for user creation with `Validate()` method returning validation errors
- **`db/`** — Data access layer (repository pattern)
  - `UserStore` interface — contract for user persistence operations
  - `MongoUserStore` — MongoDB implementation using BSON serialization
  - All operations take a `context.Context` as the first parameter
- **`api/`** — HTTP handlers
  - `UserHandler` — dependency-injected with a `UserStore` implementation
  - Returns JSON responses with appropriate HTTP status codes

**Key design**: User passwords are hashed with bcrypt before storage (never stored in plain text).

## Common Commands

```bash
# Build the application
make build

# Run (builds first)
make run

# Run all tests with verbose output
make test

# Run a single test file
go test -v ./api

# Run a specific test
go test -v -run TestNamePattern ./api

# Build Docker image
make docker

# Database seeding (if scripts/seed.go exists)
make seed
```

## Development Setup

MongoDB must be running on `localhost:27017` for local development. Connection string and database details are hardcoded in `main.go`:
- URI: `mongodb://localhost:27017`
- Database: `hotel_reservation`
- Users collection: `users`

To run MongoDB locally, use docker-compose:
```bash
docker-compose up
```

## Testing

- Test files follow Go convention: `*_test.go` (e.g., `api/user_handler_test.go`)
- Use `-count=1` flag to disable test result caching: `go test -count=1 ./...`
- Tests in `api/user_handler_test.go` validate handler behavior

## Notes for Future Development

- MongoDB ObjectID is used as the user ID (stored as `_id` in BSON, converted to hex string in JSON responses)
- All request bodies are validated at the type level before reaching handlers
- Fiber's error handler is configured to return JSON error responses
- API routes are versioned under `/api/v1/` prefix
