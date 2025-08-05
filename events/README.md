# Getting Started

This project is a secure REST API service designed to manage events such as concerts, shows, and theater performances. It includes features like JWT-based authentication, role-based access control, and comprehensive monitoring and logging.

## Prerequisites

Before running the project, ensure you have the following installed:

- Go 1.22.1 or higher
- PostgreSQL 14.18
- Docker and Docker Compose (optional, for containerized setup)

## Installation

1. Clone the repository:
```sh
git clone https://github.com/janapc/event-tickets.git
cd event-tickets/events
```

2. Copy the example environment file and configure it:
```sh
cp .env_example .env
```

3. Install dependencies:
```sh
go mod tidy
```

## Running the Application

### Local Development

1. Start the PostgreSQL database and ensure it matches the configuration in `.env`.

2. Run the application:
```sh
go run cmd/main.go
```

The API will be available at `http://localhost:3001`.

### Using Docker

1. Build and start the services using Docker Compose:
```sh
docker-compose up --build -d
```

2. Access the API at `http://localhost:3001`.

## Testing

Run the test suite to ensure everything is working correctly:
```sh
go test -v ./...
```

## Monitoring and Logging

This project includes integrated monitoring and logging tools:

- **Prometheus**: Metrics available at `http://localhost:9090`
- **Grafana**: Dashboards available at `http://localhost:3006` (default credentials: admin/admin)
- **Jaeger**: Tracing available at `http://localhost:16686`
- **Kibana**: Logs available at `http://localhost:5601`

## API Documentation

Swagger documentation is automatically generated and available at:
```
http://localhost:3001/api/index.html
```

## Development

To update the Swagger documentation after making API changes:
```sh
swag init -g cmd/main.go --output internal/infra/docs
```

For containerized development, ensure all services are running via Docker Compose.

## Author

Made by Janapc 🤘 [Get in touch!](https://www.linkedin.com/in/janaina-pedrina/)
