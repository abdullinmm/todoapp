package handlers

import (
	"encoding/json"
	"net/http"
)

// ParseLoginRequest parses request from an HTTP request body.
func ParseLoginRequest(r *http.Request) (*LoginRequest, error) {
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return &req, err
}
