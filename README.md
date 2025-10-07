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
```
docker run --name todoapp-pg
-e POSTGRES_USER=todo
-e POSTGRES_PASSWORD=secret
-e POSTGRES_DB=todoapp
-p 5432:5432 -d postgres:16
```

3. Apply migrations (choose tool: goose or golang-migrate)
for example:
```
migrate -path ./migrations -database "postgres://todo:secret@localhost:5432/todoapp?sslmode=disable" up
```

4. Add a .env file:

```
PORT=8080
JWT_SECRET=supersecretkey123
DATABASE_URL=postgres://todo:secret@localhost:5432/todoapp?sslmode=disable
```

6. Run the app:
```
go run ./cmd/todoapp
```
or use Docker Compose for both app and db:
```
docker-compose up --build
```

---

## API Example Usage

**Register:**
```
curl -X POST http://localhost:8080/register
-H "Content-Type: application/json"
-d '{"username":"demo","password":"mysecretpass"}'
```

**Login:**
```
curl http://localhost:8080/me
-H "Authorization: Bearer <JWT_token_here>"
```

**Authenticated user profile:**
```
curl http://localhost:8080/me
-H "Authorization: Bearer <JWT_token_here>"
```

---

## Testing

Run all tests with:
```
go test ./... -v
```
Coverage badge coming soon!

---

## Project structure

- `cmd/todoapp/` — entrypoint main.go, HTTP server
- `internal/auth/` — authentication logic (bcrypt/JWT)
- `internal/handlers/` — HTTP handlers and middleware
- `internal/db/` — DB access, migrations
- `migrations/` — SQL migration scripts

---

## Roadmap

- [ ] Add todos CRUD endpoints and logic
- [ ] Docker image/push steps for production use
- [ ] Add coverage badge (Codecov)
- [ ] English doc as default

---

## License

MIT © Marsel Abdullin

---

## Contacts

- Email: abdullinmm@gmail.com
- Telegram: [@abdullin_marsel](https://t.me/abdullin_marsel)
- LinkedIn: [marsel-abdullin-291238121](https://www.linkedin.com/in/marsel-abdullin-291238121/)

