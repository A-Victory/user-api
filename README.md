# User Migration Project

## Overview

This project is a user management system built with Go. It includes database migrations, a modular user package, and three main APIs for user sign-up, sign-in, and listing users.

## Requirements

- **Database**: PostgreSQL
- **Migration Tool**: `goose` (used to manage database migrations)

## Migrations

1. **Migration 1**: Create `users` table with the following fields:
    - `email` (unique, not null)
    - `password` (encrypted, not null)

2. **Migration 2**: Add new fields to the `users` table:
    - `first_name` (varchar(100))
    - `last_name` (varchar(100))

Migrations are stored in the `user/db/migration` folder and managed using the `goose` package.

## User Package

A separate `user` package is created to handle user-related operations. This package encapsulates:
- Database migrations
- Models
- Handlers
- Services
- Routes

The package is reusable and modular.

## APIs

### 1. User Sign-Up

- **Endpoint**: `/api/signup`
- **Method**: POST
- **Description**: Allows a new user to sign up by providing email, password, first name, and last name.

### 2. User Sign-In

- **Endpoint**: `/api/signin`
- **Method**: POST
- **Description**: Allows an existing user to sign in by providing email and password.

### 3. Listing Users

- **Endpoint**: `/api/get-users`
- **Method**: GET
- **Description**: Retrieves a list of all users.

## Project Structure

```
.
├── main
│   └── main.go
├── user
│   ├── db
│   │   ├── db.go
│   │   ├── setup.go
│   │   ├── migration
│   │   │   ├── 20250109201913_create_users_table.sql
│   │   │   └── 20250109202102_add_user_fields.sql
│   ├── handler
│   │   └── handlers.go
│   ├── models
│   │   └── user.go
│   ├── routes
│   │   └── routes.go
│   └── service
│       └── service.go
├── .env
├── .gitignore
├── go.mod
└── go.sum
```

## Setup

### 1. Clone the Repository
```sh
git clone <repository-url>
cd user-migration
```

### 2. Set Up Environment Variables
Create a `.env` file with the following content:
```dotenv
DB_CONFIG="user=username password=password dbname=yourdb sslmode=disable"
DB_NAME="your_database_name"
```

Replace `username`, `password`, and `yourdb` with your PostgreSQL database credentials.

### 3. Install Dependencies
Ensure you have the following installed:
- Go
- PostgreSQL
- `goose` migration tool:
  ```sh
  go install github.com/pressly/goose/v3/cmd/goose@latest
  ```

### 4. Run Database Migrations
Navigate to the `user/db/migration` folder and run migrations locally:
```sh
goose -dir ./user/db/migration postgres "user=username password=password dbname=yourdb sslmode=disable" up
```

### 5. Start the Server
Run the application:
```sh
go run main/main.go
```
Note: One starting the server, the migrations will also run automatically

## Usage

- Use tools like Postman or `curl` to interact with the APIs.
- Ensure the PostgreSQL database is running and accessible with the provided configuration.

## Example API Requests

### User Sign-Up
```bash
curl -X POST -H "Content-Type: application/json" -d '{"email":"test@example.com", "password":"password123", "first_name":"John", "last_name":"Doe"}' http://localhost:8080/api/signup
```

### User Sign-In
```bash
curl -X POST -H "Content-Type: application/json" -d '{"email":"test@example.com", "password":"password123"}' http://localhost:8080/api/signin
```

### List Users
```bash
curl -X GET http://localhost:8080/api/get-users
```
