package handlers

import (
	"encoding/json"
	"net/http"
)

// ParseRegisterRequest parses a JSON request body into a RegisterRequest struct.
func ParseRegisterRequest(r *http.Request) (*RegisterRequest, error) {
	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return &req, err
}
