# Users Service

A NestJS-based microservice for managing user-related operations in the Event Tickets system.


## Features

- RESTful API with Swagger documentation
- MongoDB integration
- JWT-based authentication
- CQRS pattern implementation
- Input validation
- CORS enabled
- Environment-based configuration

## Setup

1. Install dependencies:
```bash
npm install
```

2. Configure environment variables:
```bash
cp .env.example .env
```

Required environment variables:
- `MONGODB_URL`: MongoDB connection string
- `JWT_SECRET`: Secret key for JWT token generation
- `JWT_EXPIRES_IN`: JWT token expiration time
- `PORT`: Application port (default: 3000)
- `BASE_API_URL`: Base URL for API endpoints

3. Start the application:
```bash
npm run start:dev
```

The API documentation will be available at `http://localhost:3000/api` (or your configured base URL).

## Architecture

The project follows Clean Architecture principles with:

- **Domain Layer**: Contains core business entities and rules
- **Application Layer**: Implements use cases and orchestrates domain objects
- **Interface Layer**: Handles external communication (API, DTOs)
- **Infrastructure Layer**: Implements technical details (database, external services)

## Testing

Run the test suite:
```bash
npm test
```
