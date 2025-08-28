package main

import (
	"log"
	"net/http"

	"github.com/abdullinmm/todoapp/internal/config"
	"github.com/abdullinmm/todoapp/internal/db"
	"github.com/abdullinmm/todoapp/internal/handlers"
)

func main() {
	cfg := config.LoadConfig()

	// Do not log in the secret in Production!
	// log.Printf("jwt secret: %q", cfg.JWTSecret)

	database, err := db.InitDB(cfg.DatabaseURl)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	http.HandleFunc("/register", handlers.RegisterHandler(database))
	http.HandleFunc("/login", handlers.LoginHandler(cfg, database))
	http.Handle("/me",
		handlers.AuthMiddleware(cfg.JWTSecret, http.HandlerFunc(handlers.MeHandler)),
	)

	log.Printf("Server started on : %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
