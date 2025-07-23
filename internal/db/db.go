package db

import "database/sql"

func initDB() (*sql.DB, error) {
	dsn := "host=localhost port=5432 user=todo password=secret dbname=todoapp sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, db.Ping()
}
