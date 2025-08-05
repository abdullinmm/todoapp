package main

import (
	"log"
	"net/http"

	"github.com/abdullinmm/todoapp/internal/db"
	"github.com/abdullinmm/todoapp/internal/handlers"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Brush error to the database: %v", err)
	}
	defer database.Close()

	http.HandleFunc("/register", handlers.RegisterHandler(database))
	http.HandleFunc("/login", handlers.LoginHandler(database))

	log.Println("Server started on : 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
