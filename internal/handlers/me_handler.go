package handlers

import (
	"fmt"
	"net/http"
)

// MeHandler handles the /me endpoint
func MeHandler(w http.ResponseWriter, r *http.Request) {
	// extract userid from the context that Middleware laid
	userID := GetUserID(r)

	if userID == 0 {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"user_id": %d}`, userID)))
}
