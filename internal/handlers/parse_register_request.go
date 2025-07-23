package handlers

import (
	"encoding/json"
	"net/http"
)

// RegisterRequest represents a registration request.
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ParseRegisterRequest parses a registration request from an HTTP request.
func ParseRegisterRequest(r *http.Request) (*RegisterRequest, error) {
	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}
