package handlers

import "errors"

// ValidateRegister validates the JSON request body for the register endpoint.
func ValidateRegister(reg *RegisterRequest) error {
	if reg.Username == "" {
		return errors.New("username is required")
	}
	if reg.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
