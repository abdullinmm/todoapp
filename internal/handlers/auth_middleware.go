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
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "no token", http.StatusUnauthorized)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		userID, err := auth.ParseJWT(token, secret)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
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
