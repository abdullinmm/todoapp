package handlers

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/abdullinmm/todoapp/internal/db"

	"golang.org/x/crypto/bcrypt"
)

// LoginHandler handles login requests
func LoginHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		req, err := ParseLoginRequest(r)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		user, err := db.GetUserByUsername(database, req.Username)
		if err != nil {
			if errors.Is(err, db.ErrUserNotFound) {
				http.Error(w, "Invalid username or password", http.StatusUnauthorized)
				return
			}
			log.Println("DB error:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Comparison of the entered password with a hash
		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
		if err != nil {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		// Successful authorization
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login successful"))
	}

}
