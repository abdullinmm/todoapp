package handlers

import (
	"errors"
)

func ValidateLoginRequest(reg *LoginRequest) error {
	if len(reg.Username) < 3 || len(reg.Username) > 64 {
		return errors.New("username must be 3..64 chars")
	}
	if len(reg.Password) < 8 || len(reg.Password) > 72 {
		return errors.New("password must be 8..128 chars")
	}
	return nil
}
