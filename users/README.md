# Users Service

A robust User Management Service built with NestJS, featuring CQRS pattern, OpenTelemetry observability, and containerized deployment.

## Features

- User registration and authentication
- Role-based user management (Admin/Public)
- JWT-based authentication
- MongoDB database integration
- Comprehensive observability stack:
  - Prometheus metrics
  - Jaeger distributed tracing
  - ELK Stack logging (Elasticsearch, Logstash, Kibana)
  - Grafana dashboards
- OpenTelemetry integration
- Health checks
- API documentation with Swagger
- Docker containerization

## Tech Stack

- NestJS
- MongoDB
- JWT for authentication
- OpenTelemetry
- Docker & Docker Compose
- NGINX
- ELK Stack
- Prometheus & Grafana
- Jaeger

## Prerequisites

- Docker and Docker Compose
- Node.js 20+
- npm or yarn

## Project Structure

```
users/
├── src/
│   ├── application/     # Application layer (CQRS commands)
│   ├── domain/         # Domain layer (entities, repositories)
│   ├── infra/          # Infrastructure layer
│   └── interfaces/     # Interface layer (controllers, DTOs)
├── .docker/           # Docker configurations
└── docker-compose.yaml
```

## Environment Variables

Create a `.env.production` file with:

```env
MONGODB_URL=mongodb://example:example@mongodb:27017
PORT=3000
PREFIX=v1
JWT_SECRET=your-secret-key
JWT_EXPIRES_IN=3600
BASE_API_URL=http://localhost
OTEL_EXPORTER_OTLP_ENDPOINT=http://collector:4318
```

## Getting Started

1. Clone the repository
2. Set up environment variables
3. Start the services:

```bash
docker-compose up -d
```

The service will be available at `http://localhost/users`

## API Endpoints

- `POST /users` - Register new user
- `POST /users/token` - Generate authentication token
- `DELETE /users/:id` - Remove user
- `GET /health` - Service health check

## Monitoring & Observability

- Grafana: `http://localhost:3001`
- Kibana: `http://localhost:5601`
- Jaeger UI: `http://localhost:16686`
- Prometheus: `http://localhost:9090`

## Development

1. Install dependencies:
```bash
npm install
```

2. Run tests:
```bash
npm test
```

3. Start development server:
```bash
npm run start:dev
```
