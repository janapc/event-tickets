# Payment Service

A microservice for handling payments and transactions in an event ticketing system, built with Go and following Clean Architecture principles.


## 🚀 Features

- **Payment Processing**: Create and manage payment transactions
- **Transaction Management**: Handle payment gateway interactions with simulation
- **Event-Driven Architecture**: Kafka-based messaging for service communication
- **Email Notifications**: Multi-language email support for payment status
- **Observability**: Full observability stack with metrics, logs, and tracing
- **Database**: PostgreSQL with connection pooling
- **Testing**: Comprehensive unit tests with mocks

## 🛠️ Technology Stack

### Core Technologies
- **Go 1.24**: Primary programming language
- **Fiber v2**: Web framework for HTTP API
- **PostgreSQL 14**: Primary database
- **Kafka**: Message streaming platform

### Observability & Monitoring
- **OpenTelemetry**: Distributed tracing and metrics
- **Jaeger**: Distributed tracing UI
- **Prometheus**: Metrics collection
- **Grafana**: Metrics visualization
- **Elasticsearch + Kibana**: Log aggregation and analysis
- **Filebeat**: Log shipping

### Development & Deployment
- **Docker & Docker Compose**: Containerization
- **Testify**: Testing framework with mocks

## 📋 Prerequisites

- Docker and Docker Compose
- Go 1.24+ (for local development)

## 🚀 Quick Start

### 1. Clone and Setup

```bash
git clone <repository-url>
cd event-tickets/payments
```

### 2. Environment Configuration

```bash
# Copy environment template
cp .env_example .env.production/.env

# Configure your environment variables in .env.production or .env
```

### 3. Setup Secrets

Create the required secret files:

```bash
# Database secrets
echo "your_db_user" > .docker/secrets/postgres_user.txt
echo "your_db_password" > .docker/secrets/postgres_password.txt

# Grafana secrets
echo "admin" > .docker/secrets/grafana_user.txt
echo "your_grafana_password" > .docker/secrets/grafana_password.txt
```

### 4. Start Services

```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f payments_service
```

## 📡 API Endpoints

### Create Payment
```http
POST /payments
Content-Type: application/json

{
  "user_name": "John Doe",
  "user_email": "john@example.com",
  "event_id": "event-123",
  "event_amount": 99.99,
  "payment_token": "TOKEN123",
  "event_name": "Concert 2024",
  "event_description": "Amazing concert event",
  "event_image_url": "https://example.com/image.jpg",
  "user_language": "en"
}
```

## 🔄 Event Flow

The service participates in the following event-driven flows:

1. **Payment Creation Flow**:
   - `POST /payments` → Creates payment → Publishes `PAYMENT_CREATED`
   - Creates transaction → Publishes `TRANSACTION_CREATED`
   - Processes transaction → Publishes `TRANSACTION_SUCCEEDED/FAILED`
   - Updates payment → Publishes `PAYMENT_SUCCEEDED/FAILED`
   - Sends notification email

2. **Kafka Topics**:
   - `PAYMENT_CREATED_TOPIC`
   - `PAYMENT_SUCCEEDED_TOPIC`
   - `TRANSACTION_CREATED_TOPIC`
   - `TRANSACTION_FAILED_TOPIC`
   - `TRANSACTION_SUCCEEDED_TOPIC`

## 🧪 Testing

### Run Unit Tests
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./internal/application/payment/command/...
```
