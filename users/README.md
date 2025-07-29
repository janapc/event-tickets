# Users Service Documentation

The Users Service is a microservice built with [NestJS](https://nestjs.com/) for managing user registration, authentication, and removal. It supports JWT-based authentication and is instrumented for observability with Prometheus, Grafana, Jaeger, and the Elastic Stack.

---

## Architecture

- **NestJS**: Main application framework.
- **MongoDB**: User data storage.
- **Docker Compose**: Orchestrates all services.
- **OpenTelemetry**: Metrics and tracing.
- **Prometheus & Grafana**: Metrics collection and visualization.
- **Jaeger**: Distributed tracing.
- **Elastic Stack (Elasticsearch, Kibana, Filebeat)**: Centralized logging.
- **APM Server**: Application performance monitoring.

### Service Diagram

```
[Client]
   |
   v
[Users Service (NestJS)] <--> [MongoDB]
   |         |         |         |
   |         |         |         |
[Prometheus][Jaeger][Filebeat][APM Server]
   |         |         |         |
[Grafana] [Kibana] [Elasticsearch]
```

---

## Environment Variables

The service uses environment variables for configuration. See `.env.example` for reference.

---

## Running the Service

1. **Clone the repository** and navigate to the `event-tickets/users` directory.

2. **Set up environment variables**:
   - Copy `.env.example` to `.env.production` and adjust values as needed.

3. **Prepare Docker secrets**:
   - Ensure the `.docker/secrets/` files exist and contain the correct credentials.

4. **Start the stack**:
   ```
   docker-compose up --build
   ```

5. **Access the service**:
   - Users API: `http://localhost:3000/v1/api`
   - Swagger UI: `http://localhost:3000/v1/api`
   - Grafana: `http://localhost:3001` (default user/pass: see secrets)
   - Kibana: `http://localhost:5601`
   - Jaeger: `http://localhost:16686`
   - Prometheus: `http://localhost:9090`
---
