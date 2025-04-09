# Yeloo Clients

Yeloo Clients is a Go-based RESTful API for managing user profiles. It provides endpoints for creating, retrieving, and deleting profiles, with a PostgreSQL database as the backend.

## Features

- **CRUD Operations**: Create, retrieve, and delete user profiles.
- **Validation**: Input validation for profile data.
- **Swagger Documentation**: Auto-generated API documentation.
- **Docker Support**: Dockerized application with PostgreSQL.

## Prerequisites

- [Go](https://golang.org/dl/) 1.20 or later
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/yeloo-clients.git
   cd yeloo-clients
   ```

2. Copy the example environment file and configure it:
   ```bash
   cp example.env .env
   ```

3. Build and run the application using Docker Compose:
   ```bash
   make docker-run
   ```

4. Access the API at `http://localhost:8080`.

## API Endpoints

### Profiles

- **GET** `/profile/{id}`: Retrieve a profile by ID.
- **POST** `/profile`: Create a new profile.
- **DELETE** `/profile/{id}`: Delete a profile by ID.

### Swagger Documentation

Access the Swagger documentation at `http://localhost:8080/swagger/index.html`.

## Running Locally

1. Install dependencies:
   ```bash
   go mod download
   ```

2. Run the application:
   ```bash
   go run cmd/server/main.go
   ```

3. Run tests:
   ```bash
   make test
   ```

## Project Structure

```
├── cmd/                # Main application entry point
├── config/             # Configuration loading
├── docs/               # Swagger documentation
├── internal/           # Application logic
│   ├── controllers/    # HTTP handlers
│   ├── database/       # Database connection and migrations
│   ├── dependencies/   # Dependency injection
│   ├── models/         # Data models
│   ├── repositories/   # Database access logic
│   ├── routes/         # API routes
│   └── services/       # Business logic
├── Dockerfile          # Docker build configuration
├── docker-compose.yml  # Docker Compose configuration
├── Makefile            # Common tasks
└── README.md           # Project documentation
```

## Environment Variables

| Variable       | Description               | Default Value |
|----------------|---------------------------|---------------|
| `SERVER_PORT`  | Port for the API server   | `8080`        |
| `DB_HOST`      | Database host             | `db`          |
| `DB_PORT`      | Database port             | `5432`        |
| `DB_USER`      | Database username         | `user`        |
| `DB_PASSWORD`  | Database password         | `password`    |
| `DB_NAME`      | Database name             | `dbname`      |

## Contributing

1. Fork the repository.
2. Create a new branch:
   ```bash
   git checkout -b feature-name
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add feature"
   ```
4. Push to the branch:
   ```bash
   git push origin feature-name
   ```
5. Open a pull request.
