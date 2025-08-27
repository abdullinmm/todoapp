package handlers

import (
	"encoding/json"
	"net/http"
)

// MeResponse represents the response for the /me endpoint
type MeResponse struct {
	UserID int `json:"user_id"`
}

// MeHandler handles the /me endpoint
func MeHandler(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r)

	if userID == 0 {
		writeJSONError(w, http.StatusUnauthorized, "user_not_found", "User not found in context")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(MeResponse{UserID: userID})
}
