package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/abdullinmm/todoapp/internal/auth"
	"github.com/abdullinmm/todoapp/internal/config"
	"github.com/abdullinmm/todoapp/internal/db"
)

// LoginHandler handles login requests
func LoginHandler(cfg *config.Config, database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		req, err := ParseLoginRequest(r)
		if err != nil {
			writeJSONError(w, http.StatusBadRequest, "Invalid JSON", err.Error())
			return
		}

		err = ValidateLoginRequest(req)
		if err != nil {
			writeJSONError(w, http.StatusBadRequest, "Invalid username or password", err.Error())
			return
		}

		user, err := db.GetUserByUsername(database, req.Username)
		if err != nil {
			if errors.Is(err, db.ErrUserNotFound) {
				writeJSONError(w, http.StatusUnauthorized, "Invalid username or password", err.Error())
				return
			}
			log.Println("DB error:", err)
			writeJSONError(w, http.StatusInternalServerError, "Internal server error", err.Error())
			return
		}

		// Comparison of the entered password with a hash
		if auth.CheckPasswordHash(user.PasswordHash, req.Password) == false {
			writeJSONError(w, http.StatusUnauthorized, "Invalid username or password", "Invalid username or password")
			return
		}

		token, err := auth.GenerateJWT(user.ID, cfg.JWTSecret)
		if err != nil {
			log.Printf("JWT generation error: %v", err)
			writeJSONError(w, http.StatusInternalServerError, "Failed to generate token", err.Error())
			return
		}

		// Successful authorization
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login successful\n"))
		_ = json.NewEncoder(w).Encode(map[string]string{
			"token": token,
		})

		// fmt.Printf("w, %v", w)
	}

}
