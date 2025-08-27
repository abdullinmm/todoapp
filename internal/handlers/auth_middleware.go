package handlers

import (
	"context"
	"net/http"
	"strings"

	"github.com/abdullinmm/todoapp/internal/auth"
)

type contextKey string

const userIDKey = contextKey("userID")

// AuthMiddleware is a middleware that authenticates requests using a JWT token.
func AuthMiddleware(secret string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		// Check the availability of the header
		if authHeader == "" {
			writeJSONError(w, http.StatusUnauthorized, "missing_token", "Authorization header required")
			return
		}

		// Strict checking scheme
		if !strings.HasPrefix(authHeader, "Bearer ") {
			writeJSONError(w, http.StatusUnauthorized, "invalid_auth_header", "Authorization header must start with 'Bearer '")
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			writeJSONError(w, http.StatusUnauthorized, "empty_token", "Token cannot be empty")
			return
		}

		userID, err := auth.ParseJWT(token, secret)
		if err != nil {
			writeJSONError(w, http.StatusUnauthorized, "invalid_token", err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), userIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserID returns the user ID from the request context.
func GetUserID(r *http.Request) int {
	userID, _ := r.Context().Value(userIDKey).(int)
	return userID
}
