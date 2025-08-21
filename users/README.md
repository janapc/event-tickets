# Users Service

A microservice for user management in the event tickets system, built with NestJS and following clean architecture principles with CQRS pattern.

## 🚀 Features

- **User Management**: Registration, authentication, and deletion
- **JWT Authentication**: Secure token-based authentication
- **Clean Architecture**: Domain-driven design with CQRS pattern
- **Comprehensive Observability**: Metrics, logging, and distributed tracing
- **Health Checks**: Built-in health monitoring
- **API Documentation**: Auto-generated Swagger documentation
- **Containerized**: Docker and Docker Compose ready

### Key Patterns

- **CQRS**: Command Query Responsibility Segregation
- **Repository Pattern**: Abstract data access layer
- **Domain Entities**: Rich business objects with encapsulated logic

## 🛠️ Tech Stack

- **Framework**: NestJS
- **Language**: TypeScript
- **Database**: MongoDB with Mongoose ODM
- **Authentication**: JWT
- **Validation**: class-validator
- **Testing**: Jest
- **Documentation**: Swagger/OpenAPI
- **Containerization**: Docker

### Observability Stack

- **Metrics**: OpenTelemetry + Prometheus + Grafana
- **Tracing**: Jaeger
- **Logging**: Winston + ELK Stack (Elasticsearch, Kibana, Filebeat)
- **APM**: Elastic APM

## 🚦 Getting Started

### Prerequisites

- Node.js 23+
- Docker and Docker Compose
- MongoDB (or use the provided Docker setup)

### Environment Setup

1. Copy the environment example file and configure your environment variables:
```bash
cp .env.example .env
```

### Development

1. Install dependencies:
```bash
npm install
```

2. Start the development server:
```bash
npm run start:dev
```

The service will be available at `http://localhost:3000`

### Production with Docker

1. Create production environment file:
```bash
cp .env.example .env.production
```

2. Create Docker secrets:
```bash
mkdir -p .docker/secrets
echo "root" > .docker/secrets/mongodb_password.txt
echo "admin" > .docker/secrets/grafana_user.txt
echo "admin" > .docker/secrets/grafana_password.txt
```

3. Start all services:
```bash
docker-compose up -d
```

## 📡 API Endpoints

### Documentation

- `GET /api` - Swagger API documentation

## 🧪 Testing

Run the test suite:

```bash
# Unit tests
npm run test

# Test coverage
npm run test:cov

# E2E tests
npm run test:e2e
```

## 📊 Monitoring and Observability

The service includes comprehensive monitoring setup:

### Metrics (Prometheus + Grafana)
- **Prometheus**: `http://localhost:9090`
- **Grafana**: `http://localhost:3001` (admin/admin)

Custom metrics tracked:
- User creation count
- HTTP request count and duration
- Application performance metrics

### Tracing (Jaeger)
- **Jaeger UI**: `http://localhost:16686`
- Distributed tracing across service calls
- OpenTelemetry instrumentation

### Logging (ELK Stack)
- **Kibana**: `http://localhost:5601`
- **Elasticsearch**: `http://localhost:9200`
- Structured JSON logging
- Container log aggregation

## 🔒 Security

- **Password Hashing**: Bcrypt with salt rounds
- **JWT Tokens**: Secure token-based authentication
- **Input Validation**: class-validator for DTO validation
- **Environment Variables**: Sensitive data in environment files
