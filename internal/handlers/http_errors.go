package handlers

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse represents an error response.
type ErrorResponse struct {
	Error   string `json:"message"`
	Details string `json:"details,omitempty"`
}

func writeJSONError(w http.ResponseWriter, status int, msg string, details string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(ErrorResponse{Error: msg, Details: details})
}
