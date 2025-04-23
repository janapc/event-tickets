# Users Service

A NestJS-based microservice for managing user-related operations in the Event Tickets system. The service implements Clean Architecture principles and includes features like authentication, metrics collection, and distributed tracing.

## Features

- User registration and management
- JWT-based authentication
- Role-based access control (ADMIN, PUBLIC)
- CQRS implementation for better separation of concerns
- OpenTelemetry integration for observability
- Prometheus metrics collection
- Health checks endpoint
- RESTful API with Swagger documentation
- MongoDB integration
- Docker support for local development
- Unit tests coverage

## Technical Stack

- NestJS - Node.js framework
- MongoDB - Database
- JWT - Authentication
- OpenTelemetry - Distributed tracing
- Prometheus - Metrics collection
- Grafana - Metrics visualization
- Jaeger - Distributed tracing visualization
- Jest - Testing framework
- Swagger - API documentation

## Setup

1. Install dependencies:
```bash
npm install
```

2. Configure environment variables:
```bash
cp .env.example .env
```

Required environment variables:
- `MONGODB_URL`: MongoDB connection string
- `JWT_SECRET`: Secret key for JWT token generation
- `JWT_EXPIRES_IN`: JWT token expiration time
- `BASE_API_URL`: Base URL for API endpoints (default: v1)
- `OTEL_SERVICE_NAME`: Service name for OpenTelemetry (default: users-service)
- `PORT`: Application port (default: 3000)

3. Start the development environment:
```bash
docker-compose up -d
```

4. Start the application:
```bash
npm run start:dev
```

## API Endpoints

### Users
- `POST /users` - Register a new user
- `DELETE /users/:id` - Remove a user
- `POST /users/token` - Generate authentication token

### Health
- `GET /health` - Check service health status

## Testing

Run the test suite:
```bash
# Unit tests
npm run test

# E2E tests
npm run test:e2e

# Test coverage
npm run test:cov
```

## Docker Services

The included docker-compose.yaml provides:

- MongoDB database
- Prometheus for metrics collection
- Grafana for metrics visualization
- Jaeger for distributed tracing

Access points:
- Swagger UI: http://localhost:3000/v1/api
- Grafana: http://localhost:3001 (admin/admin)
- Prometheus: http://localhost:9090
- Jaeger UI: http://localhost:16686
