package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

// ErrInvalidToken is returned when the JWT token is invalid.
var ErrInvalidToken = errors.New("invalid JWT token")

// ParseJWT parses a JWT token and returns the user ID.
func ParseJWT(tokenString string, secret string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return 0, ErrInvalidToken
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, ErrInvalidToken
	}
	userIDf, ok := claims["user_id"].(float64)
	if !ok {
		return 0, ErrInvalidToken
	}
	return int(userIDf), nil
}
