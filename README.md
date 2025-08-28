# TodoApp (Go + JWT + Postgres + Docker)

[![Go Version](https://img.shields.io/badge/go-1.24+-blue)](https://golang.org) [![Docker Pulls](https://img.shields.io/docker/pulls/abdullinmm/todoapp)]()

## Оглавление

- [Описание](#описание)
- [Функциональность](#функциональность)
- [Требования](#требования)
- [Структура проекта](#структура-проекта)
- [Конфигурация (ENV)](#конфигурация-env)
- [Локальный запуск](#локальный-запуск)
- [Запуск через Docker Compose](#запуск-через-docker-compose)
- [Сборка и запуск Docker без Compose](#сборка-и-запуск-docker-без-compose)
- [API примеры](#api-примеры)
- [Тесты](#тесты)
- [CI/CD (опционально)](#cicd-опционально)
- [Траблшутинг](#траблшутинг)
- [Лицензия](#лицензия)

---

## Описание

- HTTP API на Go с аутентификацией по JWT.
- Слои:
	- `internal/auth` — bcrypt и JWT (HS256)
	- `internal/handlers` — хендлеры + AuthMiddleware, GetUserID
	- `internal/db` — инициализация и доступ к БД
	- `internal/config` — чтение переменных окружения
- Точка входа: `cmd/todoapp/main.go`
- Эндпоинты:
	- **POST** `/register`
	- **POST** `/login`
	- **GET** `/me` (защищённый)

## Функциональность

- Регистрация нового пользователя через `/register`.
- Генерация и выдача JWT при `/login`.
- Защищённый эндпоинт `/me` для получения информации о текущем пользователе.

## Требования

- Go 1.24+ (см. `go.mod`)
- Docker и Docker Compose
- Postgres 16+
- make (опционально)

## Структура проекта

```
todoapp/
├─ cmd/
│  └─ todoapp/
│     └─ main.go           — точка входа и HTTP-сервер
├─ internal/
│  ├─ auth/                — bcrypt и JWT (HS256)
│  ├─ handlers/            — HTTP-хендлеры + AuthMiddleware, GetUserID
│  ├─ db/                  — инициализация и доступ к БД
│  └─ config/              — чтение env
├─ migrations/             — SQL-миграции (goose/golang-migrate)
├─ Dockerfile              — multi-stage сборка
├─ docker-compose.yml      — локальное окружение
└─ .env                    — локальные переменные (не коммитить)
```


## Конфигурация (ENV)

| Переменная     | Описание                             | Пример                                                                                  |
| -------------- | ------------------------------------ | --------------------------------------------------------------------------------------- |
| `PORT`         | Порт HTTP-сервера                    | `8080`                                                                                  |
| `JWT_SECRET`   | Секрет для подписи JWT (обязательно) | `supersecretkey123`                                                                     |
| `DATABASE_URL` | URL подключения к БД                 | `postgres://todo:secret@localhost:5432/todoapp?sslmode=disable`                         |

**Пример `.env`:**
```
PORT=8080
JWT_SECRET=supersecretkey123
DATABASE_URL=postgres://todo:secret@localhost:5432/todoapp?sslmode=disable
```

## Локальный запуск

1. Установить зависимости:  
```
go mod download
```
2. Поднять Postgres (через Docker):  
```
docker run –name todoapp-pg 
-e POSTGRES_USER=todo 
-e POSTGRES_PASSWORD=secret 
-e POSTGRES_DB=todoapp 
-p 5432:5432 
-d postgres:16
```

3. Применить миграции:
- **golang-migrate**
  ```
  migrate -path ./migrations \
    -database "postgres://todo:secret@localhost:5432/todoapp?sslmode=disable" up
  ```
- **goose**
  ```
  goose -dir ./migrations \
    postgres "postgres://todo:secret@localhost:5432/todoapp?sslmode=disable" up
  ```

4. Экспорт переменных окружения (или используйте `.env`):  
```
export PORT=8080
export JWT_SECRET=supersecretkey12
export DATABASE_URL=“postgres://todo:secret@localhost:5432/todoapp?sslmode=disable”
```

5. Запустить сервер:
## Запуск через Docker Compose

1. Запустить сервисы:

```
docker-compose up –build
```

2. Описание окружения в `docker-compose.yml`:  
```
services:
app:
environment:
- PORT=8080
- JWT_SECRET=${JWT_SECRET}
- DATABASE_URL=postgres://todo:secret@db:5432/todoapp?sslmode=disable
db:
image: postgres:16
volumes:
- pg/var/lib/postgresql/data
volumes:
pg:
```

3. Миграции:
- **Вариант A**: приложение прогоняет миграции при старте.
- **Вариант B**: отдельный сервис:
  ```
  migrate:
    image: migrate/migrate
    depends_on:
      - db
    volumes:
      - ./migrations:/migrations
    entrypoint:
      - /bin/sh
      - -c
      - |
        sleep 5 &&
        migrate -path=/migrations \
          -database "postgres://todo:secret@db:5432/todoapp?sslmode=disable" up
  ```

## Сборка и запуск Docker без Compose
```
docker build -t abdullinmm/todoapp:local . docker run –rm -p 8080:8080 
-e PORT=8080 
-e JWT_SECRET=supersecretkey123 
-e DATABASE_URL=“postgres://todo:secret@host.docker.internal:5432/todoapp?sslmode=disable” 
abdullinmm/todoapp:local
```

## API примеры

- **Регистрация**  
```
curl -X POST http://localhost:8080/register 
-H “Content-Type: application/json” 
-d ‘{“username”:“demo”,“password”:“mysecretpass”}’
```

- **Логин**  
```
curl -X POST http://localhost:8080/login 
-H “Content-Type: application/json” 
-d ‘{“username”:“demo”,“password”:“mysecretpass”}’
```

- **Профиль**  
```
curl http://localhost:8080/me 
-H “Authorization: Bearer <YOUR_TOKEN>”
```
## Тесты

```
go test ./… -v
```

Рекомендуется покрыть:
- `internal/auth`: `HashPassword`/`CheckPasswordHash`, `GenerateJWT`/`ParseJWT`
- `internal/handlers`: `AuthMiddleware` (валидный/битый/просроченный токен), `MeHandler`
- `internal/db`: через моки

## CI/CD (опционально)

- Секреты: `DOCKERHUB_USERNAME`, `DOCKERHUB_TOKEN` (Write)
- Пример шагов:
	1. `go vet`
	2. `golangci-lint run`
	3. `go test`
	4. `docker build` + `docker push`

## Траблшутинг

- **invalid token** → проверьте алгоритм HS256 и значение `JWT_SECRET`.
- **ECDSA expects *ecdsa.PrivateKey** → не используйте ES256 со строковым секретом.
- **go mod requires go ≥ 1.24** → обновите образ в `Dockerfile` до `golang:1.24-alpine`.
- **docker push denied** → убедитесь, что `DOCKERHUB_USERNAME` и `DOCKERHUB_TOKEN` заданы правильно.

## Лицензия

MIT © Marsel Abdullin

