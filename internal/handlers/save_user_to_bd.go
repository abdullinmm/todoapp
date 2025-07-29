package handlers

import (
	"database/sql"
)

// SaveUserToDB saves a user to the database.
func SaveUserToDB(db *sql.DB, username string, hashedPassword string) error {
	_, err := db.Exec("INSERT INTO users (username, password_hash) VALUES ($1, $2)", username, hashedPassword)
	return err
}
