package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

// ErrInvalidToken is returned when the JWT token is invalid.
var ErrInvalidToken = errors.New("invalid JWT token")

// ParseJWT parses a JWT token and returns the user ID.
func ParseJWT(tokenString string, secret string) (int, error) {
	if secret == "" {
		return 0, errors.New("JWT secret cannot be empty")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return 0, ErrInvalidToken
	}

	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, ErrInvalidToken
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("invalid user_id in token")
	}

	return int(userIDFloat), nil
}
