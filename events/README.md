# Events Service

A secure REST API service for managing events (concerts, shows, theater) with JWT-based authentication and role-based access control.

## Features

- **Event Management**: Full CRUD operations for events
- **Authentication**: JWT-based authentication for API access
- **Role-Based Access**: Support for ADMIN and PUBLIC roles
- **API Documentation**: Swagger/OpenAPI integration
- **Database**: PostgreSQL persistence with migrations
- **Testing**: Comprehensive test suite with mocking
- **Containerization**: Docker and docker-compose support

## Technical Requirements

- Go 1.22.1
- PostgreSQL 14.18
- Docker and docker-compose (optional)

## Environment Setup

1. Copy the example environment file:
```sh
cp .env_example .env
```

2. Configure the following environment variables in `.env`:
```
PORT=:3001
DB_HOST=localhost       # Use postgres-database for Docker
DB_PORT=5432
DB_USER=root
DB_PASSWORD=root
DB_NAME=event_tickets
JWT_SECRET=secret
BASE_API_URL=http://localhost:3001
```

## Running the Application

### Local Development

1. Install dependencies:
```sh
go mod tidy
```

2. Start the server:
```sh
go run cmd/main.go
```

### Using Docker

1. Build and start services:
```sh
docker-compose up -d
```

2. The API will be available at `http://localhost:3001`

## API Endpoints

The API documentation is available at `http://localhost:3001/events/docs`

## Testing

Run the test suite:

```sh
go test -v ./...
```

## API Documentation

Swagger documentation is automatically generated and available at:

```
http://localhost:3001/docs
```

## Development

To update the Swagger documentation after making API changes:

```sh
swag init -g cmd/main.go --output internal/infra/docs
```

## Author

Made by Janapc 🤘 [Get in touch!](https://www.linkedin.com/in/janaina-pedrina/)
