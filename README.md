# REST API Implementation in Go

This project is a simple REST API implementation using Go, JWT for authentication, and the Mux router for routing.

## Features

- User registration
- User login with JWT authentication
- Protected routes that require a valid JWT token
- Input validation using the `go-playground/validator` package

## Technologies Used

- Go (Golang)
- Gorilla Mux Router
- JWT (JSON Web Tokens)
- GORM (for database interactions)
- PostgreSQL

## Getting Started

### Prerequisites

- Go installed on your machine
- PostgreSQL (or your preferred database)
- `go.mod` file for dependency management

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/IamSarojPoudel/go-rest-api.git
   cd go-rest-api
   ```

2. Install the required dependencies:

   ```bash
   go mod tidy
   ```

3. Create a `.env` file in the root directory and add your environment variables:
   ```plaintext
   SERVER_ADDRESS=your_server_address
   DATABASE_URL=your_database_url
   JWT_SECRET_KEY=your_secret_key
   ```

### Running the Application

1. Start the application:

   ```bash
   go run main.go
   ```
