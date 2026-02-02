# Go HTTP Server

A simple RESTful HTTP server built with Go, demonstrating a clean, layered architecture with a PostgreSQL database and input validation.

## Project Structure

- `main.go`: The entry point of the application, now including graceful shutdown logic.
- `server.go`: Configures the HTTP server, routes, and dependency injection.
- `handlers/`: HTTP request handlers that process incoming requests and interact with storage.
- `models/`: Data structures representing the core entities (e.g., `User`).
- `storage/`: Data persistence logic using GORM and PostgreSQL.
- `go.mod`: Go module definition and dependency management.

## Features

- **User Management**: Create and retrieve users.
- **PostgreSQL Integration**: Persistent storage using GORM.
- **Dockerized**: Fully containerized setup with Docker and Docker Compose.
- **Input Validation**: Uses `go-playground/validator` for robust request body validation.
- **Graceful Shutdown**: Handles OS signals (Ctrl+C, SIGTERM) to stop the server safely.
- **Layered Architecture**: Clear separation of concerns between HTTP logic, business models, and data storage.

## API Endpoints

### Get All Users
- **URL**: `/users`
- **Method**: `GET`
- **Success Response**: `200 OK` with JSON array of users.

### Get User by ID
- **URL**: `/users?id={id}`
- **Method**: `GET`
- **Success Response**: `200 OK` with user object.
- **Error Response**: `404 Not Found` if user doesn't exist.

### Create User
- **URL**: `/users`
- **Method**: `POST`
- **Body**:
  ```json
  {
    "name": "John Doe",
    "email": "john@example.com"
  }
  ```
- **Success Response**: `201 Created` with the created user object.

## How to Build and Run

### Prerequisites
- Docker and Docker Compose installed.

### Run with Docker (Recommended)
The easiest way to start the server and the database is using Docker Compose:

```bash
docker-compose up --build
```

This will:
1. Build the Go application container.
2. Start a PostgreSQL database instance.
3. Automatically load environment variables from the `.env` file via `docker-compose`.

### Local Setup (Optional)
If you prefer to run it outside Docker:
1. Create a database named `go_user`.
2. Configure your environment variables in a `.env` file:
   ```env
   DB_HOST=localhost
   DB_USER=your_user
   DB_PASSWORD=your_password
   DB_NAME=go_user
   DB_PORT=5432
   Auth=your_secret_key
   PORT=8080
   ```
3. Run the server:
   ```bash
   go run .
   ```

## Testing with curl

**Add a user:**
```bash
curl -X POST http://localhost:8080/users \
     -H "Content-Type: application/json" \
     -d '{"name": "Alice", "email": "alice@example.com"}'
```

**List users:**
```bash
curl http://localhost:8080/users
```
