# Clients Service

A microservice for managing clients in the event ticketing system, built with Go using Clean Architecture principles.

## 🚀 Features

### Core Functionality
- **Client Management**: Create and retrieve client information
- **Event-Driven Architecture**: Kafka-based messaging for inter-service communication
- **Payment Processing**: Handles payment success events and triggers ticket generation

### Observability & Monitoring
- **Distributed Tracing**: OpenTelemetry with Jaeger
- **Metrics**: Prometheus with Grafana dashboards
- **Centralized Logging**: ELK Stack (Elasticsearch, Kibana, Filebeat)
- **Health Checks**: Liveness and readiness endpoints

### API Documentation
- **Swagger/OpenAPI**: Auto-generated API documentation

## 🛠️ Tech Stack

- **Language**: Go 1.24.2
- **Web Framework**: Fiber v2
- **Database**: PostgreSQL 14.18
- **Message Broker**: Apache Kafka
- **Observability**: OpenTelemetry, Jaeger, Prometheus, Grafana
- **Logging**: Logrus with ELK Stack
- **Testing**: Testify, SQLMock
- **Documentation**: Swagger

## 📋 Prerequisites

- Docker and Docker Compose
- Go 1.24+ (for local development)

## 🚀 Quick Start

### Using Docker Compose (Recommended)

1. **Clone the repository**
```bash
git clone <repository-url>
cd event-tickets/clients
```

2. **Set up environment**
```bash
cp .env_example .env.production/.env
# Edit .env.production or .env with your configuration
```

3. **Create required secrets**
```bash
mkdir -p .docker/secrets
echo "root" > .docker/secrets/postgres_user.txt
echo "root" > .docker/secrets/postgres_password.txt
echo "admin" > .docker/secrets/grafana_user.txt
echo "admin" > .docker/secrets/grafana_password.txt
```

4. **Start all services**
```bash
docker-compose up -d
```

### Local Development

1. **Install dependencies**
```bash
go mod tidy
```

2. **Set up environment**
```bash
cp .env_example .env
# Edit .env with your local configuration
```

3. **Run the service**
```bash
go run cmd/main.go
```

## 🎯 Event-Driven Communication

### Consumes Events
- **PAYMENT_SUCCEEDED_TOPIC**: Processes successful payments and creates clients if needed

### Produces Events
- **CLIENT_CREATED_TOPIC**: Notifies when a new client is created
- **SEND_TICKET_TOPIC**: Triggers ticket generation for clients

## 📊 Monitoring & Observability

### Access URLs (when running via docker-compose)
- **Application**: http://localhost:3004
- **Swagger API Docs**: http://localhost:3004/api/
- **Grafana**: http://localhost:3006 (admin/admin)
- **Prometheus**: http://localhost:9090
- **Jaeger**: http://localhost:16686
- **Kibana**: http://localhost:5601

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test ./internal/application -v
```
