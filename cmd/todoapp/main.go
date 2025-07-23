package main

import (
	"log"
	"net/http"

	"github.com/abdullinmm/todoapp/internal/db"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Brush error to the database: %v", err)
	}
	defer database.Close()

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Register handler"))
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Login handler"))
	})

	log.Println("Server started on : 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	// Your code here
}
