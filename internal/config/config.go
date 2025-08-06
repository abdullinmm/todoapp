package config

import (
	"log"

	"github.com/joho/godotenv"

	"os"
)

// Config represents the configuration for the application.
type Config struct {
	JWTSecret   string
	DarabaseURl string
	Port        string
}

// LoadConfig loads the configuration from environment variables.
func LoadConfig() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		JWTSecret:   os.Getenv("JWT_SECRET"),
		DarabaseURl: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
	}

	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET is not set")
	}
	if cfg.DarabaseURl == "" {
		log.Fatal("DATABASE_URL is not set")
	}
	if cfg.Port == "" {
		cfg.Port = "8080"
	}
	return cfg
}
