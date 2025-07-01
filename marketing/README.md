# Marketing Service

A microservice dedicated to lead management and marketing operations within the event ticketing system.

## Overview

The Marketing Service is designed to track and manage leads in the event ticketing ecosystem. It handles lead creation, conversion tracking, and integrates with other services through Kafka messaging.

## Features

- Lead management (create, retrieve by email, list all)
- Lead conversion tracking
- Integration with client services via Kafka
- OpenTelemetry observability
- Containerized deployment with Docker
- Comprehensive monitoring stack (Prometheus, Grafana, Jaeger, ELK stack)

## Tech Stack

- **Framework**: NestJS
- **Database**: MongoDB
- **Messaging**: Kafka
- **Containerization**: Docker
- **Observability**: OpenTelemetry, Jaeger, Prometheus, Grafana
- **Logging**: Winston, ELK stack (Elasticsearch, Logstash, Kibana)

## Getting Started

### Prerequisites

- Docker and Docker Compose
- Node.js v23 (for local development)
- NPM

### Environment Setup

1. Copy the example environment file and update the values:

```bash
cp .env-example .env.production
```

## Running the Service

### With Docker Compose

```bash
# Make sure your environment variables are set
docker-compose up -d
```

This will start:
- The marketing service on port 3005
- MongoDB on port 27017
- Kafka on port 9094
- Jaeger UI on port 16686
- Prometheus on port 9090
- Grafana on port 3001
- Kibana on port 5601

### For Development

```bash
# Install dependencies
npm install

# Run in development mode
npm run start:dev
```

## Monitoring and Observability

- **Traces**: Access Jaeger UI at http://localhost:16686
- **Metrics**: Access Prometheus at http://localhost:9090
- **Dashboards**: Access Grafana at http://localhost:3001 (default user: admin)
- **Logs**: Access Kibana at http://localhost:5601

## Testing

```bash
# Run unit tests
npm test

# Run tests with coverage
npm run test:cov

```
