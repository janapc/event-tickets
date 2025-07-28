# Ticket Service

This service is responsible for generating and sending event tickets via email. It listens to Kafka messages and creates tickets in MongoDB, then sends an email to the user with their ticket.

## Features

- Listens to Kafka topic for ticket creation requests
- Stores tickets in MongoDB
- Sends ticket emails using configurable SMTP
- Uses NestJS CQRS pattern

## Getting Started

### Prerequisites

- Docker & Docker Compose
- Node.js (for local development)
- Kafka broker (Bitnami image used in docker-compose)
- MongoDB

### Environment Variables

Copy `.env-example` to `.env` and fill in the required values.

### Docker Compose

Start dependencies (MongoDB, Kafka):

```bash
docker-compose -f tickets/docker-compose.yaml up -d
```

Make sure to set the MongoDB password in `tickets/.docker/secrets/mongodb_password.txt`.

### Running the Service

Install dependencies:

```bash
cd tickets
npm install
```

Start the service:

```bash
npm run start:dev
```

Or build and run with Docker (you'll need to create a Dockerfile):

```bash
docker build -t ticket-service .
docker run --env-file .env --network=ticket_network ticket-service
```

### Kafka Topic

The service listens to the topic specified by `SEND_TICKET_TOPIC` (default: `SEND_TICKET_TOPIC`). Example message:

```json
{
  "messageId": "ab8e2d04-a375-40df-a9d1-1c4f7135283d",
  "email": "email@email1.com",
  "name": "test",
  "eventId": "12345",
  "eventName": "Test Event",
  "eventDescription": "Test Event",
  "eventImageUrl": "https://example.com/image.jpg",
  "language": "en"
}
```

### Testing

Run unit tests:

```bash
npm run test
```
### Useful Commands

- `npm run start:dev` - Start in development mode
- `npm run test` - Run tests
