package main

import (
	"log"

	"github.com/yourname/todoapp/internal/db"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Brush error to the database: %v", err)
	}
	defer database.Close()

	// Your code here
}
