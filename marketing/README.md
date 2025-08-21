# Marketing Service

A microservice for managing marketing leads and conversions, built with NestJS and following Domain-Driven Design (DDD) and CQRS patterns.

## 🚀 Overview

The Marketing Service is responsible for:
- Managing marketing leads (creation, retrieval, conversion tracking)
- Processing client creation events from other services
- Providing REST API endpoints for lead management
- Consuming Kafka messages for real-time event processing

## 🛠️ Technologies

- **Framework**: NestJS
- **Database**: MongoDB with Mongoose
- **Message Broker**: Apache Kafka
- **Validation**: class-validator
- **Documentation**: Swagger/OpenAPI
- **Observability**: OpenTelemetry, Prometheus, Jaeger
- **Logging**: Winston with structured logging
- **Testing**: Jest
- **Container**: Docker

## 📋 Prerequisites

- Node.js 23+
- Docker and Docker Compose
- MongoDB
- Apache Kafka

## 🚀 Getting Started

### Environment Setup

1. Copy the environment template and configure your environment variables:
```bash
cp .env-example .env
```

### Local Development

1. Install dependencies:
```bash
npm install
```

2. Start the development server:
```bash
npm run start:dev
```

3. Access the API documentation:
```
http://localhost:3005/api
```

### Docker Setup

Build and run with Docker:
```bash
docker build -t marketing-service .
docker run -p 3005:3005 --env-file .env marketing-service
```

## 🔄 Message Patterns

### Kafka Consumer

The service listens to the following Kafka topics:

#### CLIENT_CREATED_TOPIC
Processes client creation events to convert leads:
```json
{
  "email": "user@example.com"
}
```

## 🧪 Testing

Run the test suite:
```bash
# Unit tests
npm run test

# Test coverage
npm run test:cov

# End-to-end tests
npm run test:e2e
```
