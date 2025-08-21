# Events Service 🎫

A robust Go-based microservice for managing events in an event ticketing system. Built with clean architecture principles, comprehensive observability, and production-ready features.

## ✨ Features

### Core Functionality
- **Event Management**: Create, read, update, delete, and list events
- **Pagination**: Efficient event listing with configurable page sizes
- **Data Validation**: Comprehensive input validation and error handling

### Security & Authentication
- **JWT Authentication**: Token-based authentication
- **Role-Based Access Control**: Admin and public user roles
- **Protected Endpoints**: Admin-only access for CUD operations

### Observability & Monitoring
- **Distributed Tracing**: OpenTelemetry integration with Jaeger
- **Metrics Collection**: Prometheus metrics for performance monitoring
- **Structured Logging**: JSON-formatted logs with trace correlation
- **Health Checks**: Built-in health check endpoint

### API Documentation
- **OpenAPI/Swagger**: Auto-generated API documentation
- **Interactive UI**: Swagger UI for API exploration

## 🚀 Getting Started

### Prerequisites
- Go 1.24.2+
- PostgreSQL
- Docker & Docker Compose (for full stack)

### Environment Setup

1. **Clone and navigate to the project**:
```bash
git clone <repository-url>
cd event-tickets/events
```

2. **Copy environment configuration and configure your environment variables**:
```bash
cp .env_example .env
```


### Running Locally

1. **Install dependencies**:
```bash
go mod tidy
```

2. **Run database migrations**:
```bash
# Create database and run the SQL script in .docker/events.sql
psql -h localhost -U your_user -d events_db -f .docker/events.sql
```

3. **Start the service**:
```bash
go run cmd/main.go
```

The service will be available at `http://localhost:3001`

### Running with Docker

**Build and run the service**:
```bash
docker build -t events-service .
docker run -p 3001:3001 --env-file .env events-service
```

## 🧪 Testing

The project includes comprehensive unit tests with mocks:

**Run all tests**:
```bash
go test ./...
```

**Run tests with coverage**:
```bash
go test -cover ./...
```

**Run specific package tests**:
```bash
go test ./internal/application/
go test ./internal/domain/
go test ./internal/infra/database/
```
