# Gateway API

A Go-based API gateway service that manages account operations with PostgreSQL database integration.

## Features

- RESTful API endpoints for account management
- PostgreSQL database integration
- Docker containerization
- Environment-based configuration
- Database migrations support

## Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose
- PostgreSQL client (optional, for direct database access)

## Getting Started

1. Clone the repository:
```bash
git clone <repository-url>
cd gateway-api
```

2. Set up environment variables:
```bash
cp .env.example .env
# Edit .env with your configuration
```

3. Start the services using Docker Compose:
```bash
docker-compose up -d
```

4. Run database migrations:
```bash
export PATH=$PATH:$(go env GOPATH)/bin
migrate -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path migrations up
```

5. Run the application:
```bash
go run cmd/app/main.go
```

## Project Structure

```
.
├── cmd/                # Application entry points
├── internal/          # Private application code
│   ├── domain/       # Domain models
│   ├── dto/          # Data Transfer Objects
│   ├── repository/   # Database operations
│   ├── service/      # Business logic
│   └── web/          # HTTP handlers and server
├── migrations/        # Database migration files
├── docker-compose.yml # Docker configuration
└── .env              # Environment variables
```

## API Endpoints

The API documentation can be found in `test.http` file, which contains example requests for testing the endpoints.

## Development

- The project uses Go modules for dependency management
- Database migrations are handled using the `migrate` tool
- Environment variables are managed through the `.env` file
