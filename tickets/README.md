# Tickets Service

A NestJS-based microservice for handling event ticket generation and distribution. This service listens to Kafka messages, creates tickets in MongoDB, and sends email notifications to attendees.

## Features

- 🎫 **Ticket Generation**: Creates unique tickets with UUID-based passports
- 📧 **Email Notifications**: Sends HTML-formatted ticket emails to attendees
- 🌐 **Multi-language Support**: Supports English and Portuguese
- 📊 **Full Observability**: Integrated monitoring, logging, and tracing
- 🔄 **Event-Driven Architecture**: Kafka-based messaging system
- 🐳 **Containerized**: Docker and Docker Compose ready

### Tech Stack

- **Framework**: NestJS
- **Database**: MongoDB with Mongoose
- **Messaging**: Apache Kafka
- **Email**: NodeMailer with NestJS Mailer
- **Observability**:
  - OpenTelemetry for tracing
  - Prometheus for metrics
  - Grafana for visualization
  - Jaeger for distributed tracing
  - Elastic Stack (ELK) for logging

## Quick Start

### Prerequisites

- Docker and Docker Compose
- Node.js 23+ (for local development)

### Environment Setup

1. Copy the environment template:
```bash
cp .env-example .env.production/.env
## Configure your environment variables in `.env.production` or `.env`
```


3. Set up Docker secrets:
```bash
mkdir -p .docker/secrets
echo "root" > .docker/secrets/mongodb_password.txt
echo "admin" > .docker/secrets/grafana_user.txt
echo "your-grafana-password" > .docker/secrets/grafana_password.txt
```

### Running with Docker Compose

```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f tickets_service

# Stop services
docker-compose down
```

## Usage

### Sending Ticket Creation Messages

The service listens to the `SEND_TICKET_TOPIC` Kafka topic. Send messages in this format:

```json
{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "eventId": "event-123",
  "eventName": "Amazing Conference 2024",
  "eventDescription": "The best tech conference of the year",
  "eventImageUrl": "https://example.com/event-image.jpg",
  "language": "en"
}
```

### Example using Kafka CLI

```bash
# Access Kafka container
docker exec -it kafka bash

# Send a test message
kafka-console-producer.sh --bootstrap-server localhost:9092 --topic SEND_TICKET_TOPIC
# Paste the JSON above and press Enter
```

## Local Development

### Setup

```bash
# Install dependencies
npm install

# Start local dependencies (MongoDB, Kafka, etc.)
docker-compose up -d mongodb kafka kafka-init-topics

# Set up local environment
cp .env-example .env
# Configure .env with local settings
```

### Running Locally

```bash
# Development mode with hot reload
npm run start:dev

# Production mode
npm run build
npm run start:prod
```

### Testing

```bash
# Unit tests
npm run test

# Test coverage
npm run test:cov

# E2E tests
npm run test:e2e
```

## Monitoring and Observability

### Accessing Dashboards

- **Grafana**: http://localhost:3001 (admin/your-password)
- **Jaeger**: http://localhost:16686
- **Kibana**: http://localhost:5601
- **Prometheus**: http://localhost:9090

## API Documentation

### Message Patterns

#### SEND_TICKET_TOPIC
Creates a new ticket and sends email notification.

**Payload:**
```typescript
{
  name: string;           // Attendee name
  email: string;          // Attendee email
  eventId: string;        // Unique event identifier
  eventName: string;      // Event display name
  eventDescription: string; // Event description
  eventImageUrl: string;  // Event image URL
  language: string;       // 'en' | 'pt'
}
```
