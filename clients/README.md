# Event Tickets - Clients Service

This project is a microservice for managing clients in an event ticketing system. It provides APIs to register new clients, retrieve client information by email, and integrates with Kafka for messaging.

## Features

- **Client Management**: Create and retrieve client information.
- **Kafka Integration**: Processes messages from Kafka topics and produces messages to other topics.
- **PostgreSQL Database**: Stores client data.
- **Swagger Documentation**: API documentation available at `/clients/docs`.

## Prerequisites

Before running the project, ensure you have the following installed:

- Docker and Docker Compose
- Go (version 1.22.1 or later)
- PostgreSQL (if running without Docker)
- Kafka (if running without Docker)

## Environment Variables

The project uses environment variables for configuration. You can find an example in the `.env_example` file. Copy it to `.env` and update the values as needed:

```bash
cp clients/.env_example clients/.env
```

### Key Environment Variables

- `PORT`: The port on which the service will run (default: `3004`).
- `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`: Database connection details.
- `KAFKA_BROKERS`: Kafka broker addresses.
- `SUCCESS_PAYMENT_TOPIC`, `CLIENT_CREATED_TOPIC`, `SEND_TICKET_TOPIC`: Kafka topic names.

## Running the Project

### Using Docker Compose

1. Build and start the services:

   ```bash
   docker-compose up --build
   ```

2. The service will be available at `http://localhost:3004`.

3. Swagger documentation can be accessed at `http://localhost:3004/clients/docs`.

### Running Locally

1. Start PostgreSQL and Kafka services.

2. Set up the database:

   - Create a database named `event_tickets`.
   - Run the SQL script located at `clients/.docker/clients.sql` to create the `clients` table.

3. Export the required environment variables or use the `.env` file.

4. Run the application:

   ```bash
   go run clients/cmd/main.go
   ```

5. The service will be available at `http://localhost:3004`.

### Testing

The project includes unit tests for the application logic and database interactions. To run the tests:

```bash
go test ./...
```

## Kafka Topics

The service interacts with the following Kafka topics:

- **`SUCCESS_PAYMENT_TOPIC`**: Consumes messages about successful payments.
- **`CLIENT_CREATED_TOPIC`**: Produces messages when a new client is created.
- **`SEND_TICKET_TOPIC`**: Produces messages to send tickets to clients.

## API Endpoints

### Health Check

- **Liveness**: `GET /clients/healthcheck/live`
- **Readiness**: `GET /clients/healthcheck/ready`

### Client Management

- **Get Client by Email**: `GET /clients?email={email}`
- **Create Client**: `POST /clients`

### Swagger Documentation

Access the API documentation at `/clients/docs`.

## Author
Made by Janapc 🤘 [Get in touch](https://www.linkedin.com/in/janaina-pedrina/)!
