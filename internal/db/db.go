package db

import (
	"database/sql"

	_ "github.com/lib/pq"

	"fmt"
)

// InitDB initializes the database connection and returns a pointer to the database object.
func InitDB(DatabaseURL string) (*sql.DB, error) {
	//dsn := "host=localhost port=5432 user=todo password=secret dbname=todoapp sslmode=disable"
	//dsn = fmt.Sprintf("%s", dsn)
	db, err := sql.Open("postgres", DatabaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Connected to Postgres successfully")
	return db, db.Ping()
}
