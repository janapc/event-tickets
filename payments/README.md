# Event Tickets Payment Service

## Overview

This service is part of an event ticketing system, responsible for handling the payment processing workflow. Built using Go and following Domain-Driven Design principles, it manages the lifecycle of payments and transactions for event ticket purchases.

## Architecture

The application is structured using a clean architecture approach:

- **Domain Layer**: Contains the core business logic and entities
- **Application Layer**: Implements use cases through command handlers
- **Interfaces Layer**: Handles external communication (HTTP API, messaging)
- **Infrastructure Layer**: Provides implementations for external dependencies

## Key Features

- Payment processing for event tickets
- Transaction management with status tracking
- Asynchronous event-driven communication using Kafka
- Email notifications for payment status changes
- Support for internationalization (i18n) for user communications

## Technical Stack

- **Language**: Go 1.24
- **Web Framework**: Fiber
- **Database**: PostgreSQL
- **Messaging**: Kafka
- **Email**: Mail service via SMTP
- **Configuration**: Environment variables via godotenv

## Getting Started

### Prerequisites

- Go 1.24+
- Docker and Docker Compose
- PostgreSQL
- Kafka

### Environment Setup

1. Clone the repository
2. Copy `.env_example` to `.env` and configure your environment variables
3. Start the required services using Docker Compose:

```bash
docker-compose up -d
```

### Running the Application

```bash
go run cmd/main.go
```

## API Endpoints

- `POST /payments` - Create a new payment

## Event Flow

1. Payment is created via the API
2. Payment creation event is published to Kafka
3. Transaction service consumes the event and creates a transaction
4. Transaction processing occurs (simulated gateway)
5. Transaction result events are published (success/failure)
6. Payment status is updated based on transaction result
7. Email notifications are sent to users

## Database Schema

The service uses two primary tables:

- `payments` - Stores payment information
- `transactions` - Stores transaction information linked to payments

## Development

### Testing

Run the tests with:

```bash
go test ./...
```
