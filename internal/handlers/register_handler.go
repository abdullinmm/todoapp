package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/abdullinmm/todoapp/internal/auth"
	"github.com/abdullinmm/todoapp/internal/db"
)

// RegisterHandler handles registration requests.
func RegisterHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		log.Println("→ registerHandler вызван")
		reg, err := ParseRegisterRequest(r)
		if err != nil {
			writeJSONError(w, http.StatusBadRequest, "", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Printf("→ получен JSON: username=%q password=%q", reg.Username, reg.Password)

		err = ValidateRegister(reg)
		if err != nil {
			writeJSONError(w, http.StatusBadRequest, "", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		hashedPassword, err := auth.HashPassword(reg.Password)
		if err != nil {
			writeJSONError(w, http.StatusInternalServerError, "", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("→ пароль захеширован: %s", hashedPassword[:10]) // первые 10 символов

		log.Println("→ пытаемся вставить пользователя в БД")
		_, err = db.CreateUser(database, reg.Username, hashedPassword)
		if err != nil {
			log.Printf("❌ ошибка при вставке: %v", err) // выброси ошибку прямо в лог
			http.Error(w, "DB insert error", 500)
			return
		}

		log.Println("✅ пользователь успешно создан")

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Register handler"))
	}
}
