package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

// Config holds application settings.
type Config struct {
    Port string
}

// LoadConfig loads configuration from .env or system environment.
func LoadConfig() *Config {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using system environment variables")
    }
   	port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // default port
    }
    return &Config{
        Port: port,
    }
}

