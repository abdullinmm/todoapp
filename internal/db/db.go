package db

import (
	"database/sql"
	"fmt"
)

func initDB() (*sql.DB, error) {
	dsn := "host=localhost port=5432 user=todo password=secret dbname=todoapp sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Connected to Postgres successfully")
	return db, db.Ping()
}
