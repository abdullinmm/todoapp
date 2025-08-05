package db

import (
	"database/sql"
	"errors"

	"github.com/abdullinmm/todoapp/internal/models"
)

// ErrUserNotFound is returned when a user is not found.
var ErrUserNotFound = errors.New("user not found")

// GetUserByUsername retrieves a user by their username.
func GetUserByUsername(db *sql.DB, username string) (*models.User, error) {
	var user models.User
	err := db.QueryRow(
		"SELECT id, username, password_hash FROM users WHERE username = $1", username,
	).Scan(&user.ID, &user.Username, &user.PasswordHash)

	if err == sql.ErrNoRows {
		return &user, ErrUserNotFound
	}

	return &user, err
}

// CreateUser creates a new user.
func CreateUser(db *sql.DB, username, password string) (*models.User, error) {
	var user models.User
	err := db.QueryRow(
		"INSERT INTO users (username, password_hash) VALUES ($1, $2) RETURNING id",
		username, password,
	).Scan(&user.ID)

	if err != nil {
		return &user, err
	}

	return &user, err
}

// SaveUserToDB saves a user to the database.
func SaveUserToDB(db *sql.DB, username string, hashedPassword string) error {
	_, err := db.Exec("INSERT INTO users (username, password_hash) VALUES ($1, $2)", username, hashedPassword)
	return err
}
