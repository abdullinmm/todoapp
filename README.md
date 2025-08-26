Проект: TodoApp (Go + JWT + Postgres + Docker)
Описание
	•	HTTP API на Go с аутентификацией по JWT.
	•	Слои: internal/auth, internal/handlers, internal/db, internal/config.
	•	Точки входа: cmd/todoapp/main.go.
	•	Эндпоинты: POST /register, POST /login, GET /me (защищённый).
Требования
	•	Go 1.24+ (см. go.mod)
	•	Docker и Docker Compose
	•	Postgres 16+
	•	make (опционально)
Структура
	•	cmd/todoapp/main.go — сервер.
	•	internal/auth — bcrypt и JWT (HS256).
	•	internal/handlers — хендлеры + AuthMiddleware, GetUserID.
	•	internal/db — инициализация и доступ к БД.
	•	internal/config — чтение env.
	•	migrations — SQL миграции (goose/golang-migrate).
	•	Dockerfile — multi-stage.
	•	docker-compose.yml — локальная среда.
	•	.env — локальные переменные (не коммитить секреты).
Конфигурация (ENV)
	•	PORT — порт HTTP сервера, по умолчанию 8080.
	•	JWT_SECRET — секрет для подписи JWT (обязателен).
	•	DATABASE_URL — строка подключения к БД: postgres://USER:PASSWORD@HOST:PORT/DB?sslmode=disable
Пример .env PORT=8080 JWT_SECRET=supersecretkey123 DATABASE_URL=postgres://todo:secret@localhost:5432/todoapp?sslmode=disable
Локальный запуск
	1.	Установить зависимости: go mod download
	2.	Поднять Postgres (через docker): docker run –name todoapp-pg -e POSTGRES_USER=todo -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=todoapp -p 5432:5432 -d postgres:16
	3.	Применить миграции:
	•	golang-migrate (если используется): migrate -path ./migrations -database ‘postgres://todo:secret@localhost:5432/todoapp?sslmode=disable’ up
	•	goose (если используется): goose -dir ./migrations postgres ‘postgres://todo:secret@localhost:5432/todoapp?sslmode=disable’ up
	4.	Экспорт ENV (или .env, если подхватывается автоматически): export PORT=8080 export JWT_SECRET=supersecretkey123 export DATABASE_URL=‘postgres://todo:secret@localhost:5432/todoapp?sslmode=disable’
	5.	Запуск сервера: go run ./cmd/todoapp/main.go
Запуск через Docker Compose
	1.	Запуск: docker-compose up –build
	2.	Окружение в compose:
	•	app: PORT=8080, JWT_SECRET, DATABASE_URL=postgres://todo:secret@db:5432/todoapp?sslmode=disable
	•	db: Postgres 16, volume pgdata
	3.	Миграции:
	•	Вариант А: приложение прогоняет миграции на старте.
	•	Вариант Б: отдельный сервис: migrate: image: migrate/migrate depends_on: db volumes:
	•	./migrations:/migrations entrypoint: ”/bin/sh”,”-c”,“sleep 5; migrate -path=/migrations -database ‘postgres://todo:secret@db:5432/todoapp?sslmode=disable’ up”
Сборка и запуск Docker без compose docker build -t abdullinmm/todoapp:local . docker run –rm -p 8080:8080 
-e PORT=8080 
-e JWT_SECRET=supersecretkey123 
-e DATABASE_URL=‘postgres://todo:secret@host.docker.internal:5432/todoapp?sslmode=disable’ 
abdullinmm/todoapp:local
API примеры Регистрация: curl -X POST http://localhost:8080/register 
-H “Content-Type: application/json” 
-d ‘{“username”:“demo”,“password”:“mysecretpass”}’ Логин: curl -X POST http://localhost:8080/login 
-H “Content-Type: application/json” 
-d ‘{“username”:“demo”,“password”:“mysecretpass”}’ Профиль: curl http://localhost:8080/me 
-H “Authorization: Bearer ”
Тесты Запуск: go test ./… -v Рекомендуется покрыть:
	•	internal/auth: HashPassword/CheckPasswordHash, GenerateJWT/ParseJWT
	•	internal/handlers: AuthMiddleware (валид/битый/просроченный токен), MeHandler
	•	internal/db: через моки
CI/CD (опционально)
	•	Секреты: DOCKERHUB_USERNAME, DOCKERHUB_TOKEN (Read/Write).
	•	Шаги: go vet, golangci-lint, go test, docker build/push.
Траблшутинг
	•	invalid token → проверь HS256 и JWT_SECRET.
	•	ECDSA expects *ecdsa.PrivateKey → не используй ES256 с строковым секретом.
	•	go mod requires go >= 1.24 → обнови образ в Dockerfile до golang:1.24-alpine.
	•	docker push denied → подставь свой Docker Hub username и токен с write-правами.
