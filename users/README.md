# Users Service

A NestJS-based microservice for managing user-related operations in the Event Tickets system. The service implements Clean Architecture principles and includes features like authentication, metrics collection, and distributed tracing.

## Technical Stack

- **Framework**: NestJS
- **Database**: MongoDB with Mongoose
- **Authentication**: JWT
- **API Documentation**: Swagger/OpenAPI
- **Observability**:
  - OpenTelemetry for tracing
  - Prometheus for metrics
  - Grafana for visualization
  - Jaeger for distributed tracing
- **Testing**: Jest
- **Infrastructure**:
  - Docker & Docker Compose
  - Nginx
- **Other Tools**:
  - Class Validator for DTO validation
  - CQRS for command handling

## API Endpoints

### Users
- `POST v1/users` - Register a new user
- `DELETE v1/users/:id` - Remove a user
- `POST v1/users/token` - Generate authentication token

### Health & Documentation
- `GET /health` - Service health check
- `GET /v1/api` - Swagger documentation

## Setup & Configuration

1. Install dependencies:
```bash
npm install
```

2. Configure environment variables:
```bash
cp .env.example .env
```

### Setting Up Secrets

1. Create the secrets directory:
```bash
mkdir -p .docker/secrets
```

2. Create secret files:
```bash
echo "your_mongodb_username" > .docker/secrets/mongodb_username ## mongoDB root username
echo "your_mongodb_password" > .docker/secrets/mongodb_password ## mongoDB root password
echo "your_grafana_username" > .docker/secrets/grafana_user ## grafana admin username
echo "your_grafana_password" > .docker/secrets/grafana_password ## grafana admin password
```

3. Ensure proper permissions:
```bash
chmod 600 .docker/secrets/*
```

### Required Environment Variables:
- `MONGODB_URL`: MongoDB connection string
- `JWT_SECRET`: Secret key for JWT tokens
- `JWT_EXPIRES_IN`: JWT token expiration time
- `BASE_API_URL`: Base API URL (default: v1)
- `OTEL_SERVICE_NAME`: OpenTelemetry service name
- `PORT`: Application port (default: 3000)
- `PREFIX`: API prefix for versioning


## Development

Start the development environment:
```bash
# Start infrastructure services
docker-compose --env-file .env.production up -d

# Start application in development mode
npm run start:dev
```

## Testing

```bash
# Unit tests
npm run test

# Test coverage
npm run test:cov
```

## Docker Services

The included docker-compose.yaml provides:

- **Nginx**: Reverse proxy (ports: 80, 443)
- **MongoDB**: Database (port: 27017)
- **Prometheus**: Metrics collection (port: 9090)
- **Grafana**: Metrics visualization (port: 3001)
- **Jaeger**: Distributed tracing (port: 16686)

### Access Points:
- Application: http://localhost/users
- Swagger UI: http://localhost/users/api
- Grafana: http://localhost:3001
- Prometheus: http://localhost:9090
- Jaeger UI: http://localhost:16686
