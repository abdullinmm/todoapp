package handlers

import (
	"encoding/json"
	"net/http"
)

// MeHandler handles the /me endpoint
func MeHandler(w http.ResponseWriter, r *http.Request) {
	// extract userid from the context that Middleware laid
	userID := GetUserID(r)

	if userID == 0 {
		writeJSONError(w, http.StatusUnauthorized, "User not found in context", "")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]int{
		"user_id": userID,
	})
}
