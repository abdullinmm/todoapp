package handlers

import "errors"

// ValidateJSON validates the JSON request body for the register endpoint.
func ValidateJSON(reg *RegisterRequest) error {
	if reg.Username == "" {
		return errors.New("username is required")
	}
	if reg.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
