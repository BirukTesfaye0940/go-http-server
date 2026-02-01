# Go HTTP Server

A simple RESTful HTTP server built with Go, demonstrating a clean, layered architecture with in-memory storage.

## Project Structure

- `main.go`: The entry point of the application.
- `server.go`: Configures the HTTP server, routes, and dependency injection.
- `handlers/`: HTTP request handlers that process incoming requests and interact with storage.
- `models/`: Data structures representing the core entities (e.g., `User`).
- `storage/`: Data persistence logic. Currently uses an in-memory repository for simplicity.
- `go.mod`: Go module definition and dependency management.

## Features

- **User Management**: Create and retrieve users.
- **In-Memory Storage**: Thread-safe storage using sync primitives.
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
- Go 1.18 or higher installed.

### Run the Server
Use the `go run` command in the root directory to start the server:
```bash
go run .
```
The server will start on `http://localhost:8080`.

### Build from Source
To compile the project into a binary:
```bash
go build -o server .
./server
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
