# Go REST API

This is a RESTful API built with Go (Golang) using the Gin framework, designed to provide a scalable and efficient backend service for managing books. The project follows clean architecture principles, separating concerns into layers such as handlers, services, repositories, and models.

## Features

- RESTful endpoints for book management
- Middleware for authentication and authorization (planned)
- Database migrations for schema management (planned)
- JSON response handling
- Modular structure for easy maintenance and extension

## Project Structure

- `cmd/api/`: Entry point of the application
- `internal/`: Internal application code
  - `config/`: Configuration management
  - `handler/`: HTTP request handlers
  - `middleware/`: Custom middleware
  - `model/`: Data models
  - `repository/`: Data access layer
  - `router/`: Route definitions
  - `service/`: Business logic layer
- `migrations/`: Database migration scripts
- `pkg/`: Shared packages
  - `response/`: JSON response utilities

## Getting Started

1. Clone the repository
2. Install dependencies: `go mod tidy`
3. Start the server: `go run cmd/api/main.go`
4. The API will be available at `http://localhost:8080`

## API Endpoints

- `GET /books`: Retrieve all books
- `POST /books`: Create a new book

Example request for POST:
```json
{
  "id": "3",
  "title": "New Book",
  "author": "Author Name"
}
```

## Technologies Used

- Go 1.24
- Gin web framework
- JSON for data exchange

For more details, refer to the code and configuration files.