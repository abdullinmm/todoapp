# TodoApp

![CI](https://github.com/abdullinmm/todoapp/actions/workflows/go.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.24-blue)
![License](https://img.shields.io/badge/license-MIT-green)

A simple yet robust backend REST API for user authentication and todo management. Built with Go, PostgreSQL, Docker, JWT, and automated CI/CD on GitHub Actions.

---

## Features

- User registration, login, JWT authentication/authorization
- Secure password hashing (bcrypt, HS256 JWT)
- CRUD operations for todos (expand as needed)
- Protected user profile endpoint (`/me`)
- Modular project structure, ENV config support
- Docker & Docker Compose for local development
- Automated tests and build (GitHub Actions)

---

## Quick Start
1. Clone the repo and install deps
go mod download

2. Start Postgres (Docker):
- docker run --name todoapp-pg
-e POSTGRES_USER=todo
-e POSTGRES_PASSWORD=secret
-e POSTGRES_DB=todoapp
-p 5432:5432 -d postgres:16

3. Apply migrations (choose tool: goose or golang-migrate)
for example:
migrate -path ./migrations -database "postgres://todo:secret@localhost:5432/todoapp?sslmode=disable" up

4. Add a .env file:
PORT=8080
JWT_SECRET=supersecretkey123
DATABASE_URL=postgres://todo:secret@localhost:5432/todoapp?sslmode=disable

5. Run the app:
go run ./cmd/todoapp

or use Docker Compose for both app and db:
docker-compose up --build

---

## API Example Usage

**Register:**

