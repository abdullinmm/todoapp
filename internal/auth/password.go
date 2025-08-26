package auth

import (
	"golang.org/x/crypto/bcrypt"
)

// Hashpassword drives a password using BCRYPT.
// Entrance: raw password.
// Exit: line-hesh and error.
func HashPassword(raw string) (string, error) {
	// bcrypt.generatefrompassword itself generates Salt inside
	// 12 - good balance of performance and security
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 12)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// Checkpasswordhash compares a raw password with hash.
// Returns True if the password is suitable, otherwise false.
func CheckPasswordHash(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
