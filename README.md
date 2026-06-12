# Golang CRUD API with PostgreSQL

A simple RESTful CRUD API built with Golang, Gin Framework, GORM, and PostgreSQL.

## Features

* Create User
* Get All Users
* Get Single User
* Update User
* Delete User
* PostgreSQL Database Integration
* GORM ORM
* UUID Primary Key
* Duplicate Email Validation
* Environment Variable Configuration

## Tech Stack

* Golang
* Gin Framework
* GORM
* PostgreSQL
* UUID
* Godotenv

## Project Structure

```bash
golang-crud-api/
│
├── cmd/
│   └── main.go
│
├── config/
│   └── database.go
│
├── controllers/
│   └── user_controller.go
│
├── models/
│   └── user.go
│
├── routes/
│   └── routes.go
│
├── .env
├── go.mod
└── README.md
```

## Installation

### Clone Repository

```bash
git clone https://github.com/your-username/golang-crud-api.git

cd golang-crud-api
```

### Install Dependencies

```bash
go mod tidy
```

### Create PostgreSQL Database

```sql
CREATE DATABASE golang_crud;
```

### Configure Environment Variables

Create a `.env` file in the project root.

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=golang_crud
```

## Run Project

```bash
go run cmd/main.go
```

Server will start on:

```text
http://localhost:8080
```

---

## API Endpoints

### Create User

**POST**

```http
/api/users
```

Request Body

```json
{
  "name": "Sekhar",
  "email": "sekhar@gmail.com"
}
```

---

### Get All Users

**GET**

```http
/api/users
```

---

### Get Single User

**GET**

```http
/api/users/{id}
```

---

### Update User

**PUT**

```http
/api/users/{id}
```

Request Body

```json
{
  "name": "Sekhar Chaudhary",
  "email": "sekhar@gmail.com"
}
```

---

### Delete User

**DELETE**

```http
/api/users/{id}
```

---

## Validation

* Email must be unique.
* Duplicate records are not allowed.
* Invalid requests return proper HTTP status codes.

## HTTP Status Codes

| Code | Description           |
| ---- | --------------------- |
| 200  | Success               |
| 201  | Created               |
| 400  | Bad Request           |
| 404  | Not Found             |
| 409  | Conflict              |
| 500  | Internal Server Error |

## Future Improvements

* JWT Authentication
* Role Based Access Control (RBAC)
* Refresh Token
* Docker Support
* Swagger Documentation
* Unit Testing
* Redis Cache
* Clean Architecture

## Author

Sekhar Chaudhary

Web Application Developer

Tech Stack: Golang, Laravel, Node.js, React, Vue, Angular, PostgreSQL
