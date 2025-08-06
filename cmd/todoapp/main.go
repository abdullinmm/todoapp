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
	log.Printf("jwt secret: %q", cfg.JWTSecret)
	database, err := db.InitDB(cfg.DarabaseURl)
	if err != nil {
		log.Fatalf("Brush error to the database: %v", err)
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
