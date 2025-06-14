# Clients Service

The Clients Service is a microservice designed to manage client information and handle messaging related to client creation and ticket processing. It is built using Go and integrates with various tools and services for telemetry, logging, messaging, and database management.

## Features

- **Client Management**: Create and retrieve client information.
- **Kafka Integration**: Consume and produce messages for client creation and ticket processing.
- **Telemetry**: OpenTelemetry integration for tracing and metrics.
- **Logging**: Structured logging using Logrus.
- **Database**: PostgreSQL for persistent storage.
- **Health Checks**: Liveness and readiness endpoints.
- **Swagger Documentation**: API documentation for easy integration.

## Architecture

The service follows a clean architecture pattern, separating the domain, application, and infrastructure layers. It uses the following technologies:

- **Go**: The programming language for the service.
- **PostgreSQL**: Database for storing client information.
- **Kafka**: Messaging system for event-driven communication.
- **OpenTelemetry**: For distributed tracing and metrics.
- **Fiber**: Web framework for building the API.
- **Swagger**: API documentation.

## Endpoints

### Health Check
- **GET** `/clients/healthcheck/live`: Liveness endpoint.
- **GET** `/clients/healthcheck/ready`: Readiness endpoint.

### Client Management
- **GET** `/clients`: Retrieve client information by email.
- **POST** `/clients`: Create a new client.

### API Documentation
- **GET** `/clients/docs/*`: Swagger documentation.

## Setup

### Prerequisites
- Docker
- Docker Compose

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/janapc/event-tickets.git
   cd event-tickets/clients
   ```
2. Create a `.env` file based on `.env_example` and fill in the required values.
3. Start the services using Docker Compose:
   ```bash
   docker-compose up --build
   ```
4. Access the service at `http://localhost:3004`.

## Telemetry and Monitoring

The service integrates with OpenTelemetry for tracing and metrics. It also includes monitoring tools like Prometheus, Grafana, Elasticsearch, and Kibana.

### Prometheus
- Access Prometheus at `http://localhost:9090`.

### Grafana
- Access Grafana at `http://localhost:3006`.
- Default credentials:
  - Username: `admin`
  - Password: `admin`

### Kibana
- Access Kibana at `http://localhost:5601`.

### Jaeger
- Access Jaeger at `http://localhost:16686`.

## Testing

The service includes unit tests for the application and infrastructure layers. Run the tests using:
```bash
go test ./...
```

## Documentation

Swagger documentation is available at `http://localhost:3004/clients/docs`.

```
