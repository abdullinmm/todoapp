package handlers

// LoginRequest represents a login request.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RegisterRequest represents a registration request.
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
