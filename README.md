# 🎵 Clean Architecture API

A RESTful API for a music store system built with Go and the Fiber framework, following the Clean Architecture principles. This project provides a structured, scalable backend with user management features, JWT-based authentication, PostgreSQL integration, and environment-based configuration management.

---

## 🚀 Features

- Clean Architecture with separation of concerns
- Fiber-based web server (lightweight and high-performance)
- Environment configuration via Viper
- JWT authentication and authorization
- User role management (`customer`)
- Secure password handling with bcrypt
- PostgreSQL integration using GORM
- Graceful error handling and middleware integration
- CLI support via Cobra

---

## 📁 Project Structure
```bash
Clean Architecture API/
├── cmd/ # CLI commands (main entry via Cobra)
├── config/ # Application configuration modules
├── internal/
│ ├── application/ # Business logic (UseCases)
│ ├── domain/ # Domain entities and interfaces
│ └── adapter/ # Infrastructure (DB, external services)
├── rest/
│ ├── api/ # HTTP routing and handlers
│ └── middleware/ # Middleware logic (auth, recovery)
├── go.mod / go.sum # Go dependencies
└── main.go # Application entry point

```
---

## ⚙️ Getting Started

### Prerequisites

- Go 1.21 or later
- PostgreSQL installed and running
- Git

### 1. Clone the Repository

```bash
git clone https://github.com/mzfarshad/clean-architecture-api
cd clean-architecture-api

```

### 2. Set Up Environment Variables

Create a .env file in the root directory:
```bash
APP_NAME=music_store_api
APP_HOST=localhost
APP_PORT=8080
APP_ENV=local
APP_DEBUG=true

DB_DSN=host=localhost user=postgres password=yourpassword dbname=music_store port=5432 sslmode=disable

JWT_ACCESS_SECRET=your_jwt_secret
JWT_ACCESS_TTL=15m
```
Replace yourpassword and your_jwt_secret with your actual credentials.

### 3. Install Dependencies
```bash
go mod tidy
```

### 4. Run the Server
```bash
go run main.go serve
```
Or build and run:
```bash
go build -o app .
./app serve
```

## 📡 API Endpoints

Base path: /api/v1

### Public

- POST /auth/signup – Register new user

- POST /auth/login – Authenticate and receive JWT token

### Protected (requires Authorization: Bearer <token>)

- PUT /user/profile – Update customer name

## 🔐 Authentication

Uses JWT for secure user sessions. You must include the access token in the request headers:

```bash
Authorization: Bearer <your-access-token>
```

##  📐 Architecture Layers

This project is built based on Clean Architecture principles:

- Domain Layer: Entities, repositories, use case interfaces

- Application Layer: Business logic implementation (UseCases)

- Infrastructure Layer: GORM repository and database access

- Delivery Layer: HTTP routes and handlers via Fiber


## 🚀 Tech Stack


| Category       | Technology / Tool              | Description                                 |
|----------------|-------------------------------|---------------------------------------------|
| Language       | Go (Golang)                   | Main backend language                       |
| Framework      | Fiber                         | Fast HTTP web framework for Go              |
| Architecture   | Clean Architecture            | Layered, decoupled structure                |
| Auth           | JWT                           | JSON Web Tokens for authentication          |
| Database       | PostgreSQL                    | Relational database                         |
| ORM            | GORM                          | ORM for interacting with PostgreSQL         |
| Config         | Viper + godotenv              | Config & environment variable management    |
| CLI            | Cobra                         | CLI commands (e.g., `serve`)                |
| Dependency Mgmt| Go Modules (`go.mod`)         | Dependency tracking                         |
| Validation     | go-playground/validator       | Struct validation                           |
| Logging        | Fiber Logger                  | HTTP request logging                        |



## 🛠️ Testing (Postman)

After starting the API server, you can test endpoints using Postman or cURL.

Example signup request:
```bash
POST /api/v1/auth/signup
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "123456",
}
```
## 📦 Future Improvements

- Add JWT Refresh Token support

- Role-based access control (Admin, Staff, etc.)

- Docker support for easier deployment

- Unit and integration tests

- Swagger documentation for API







