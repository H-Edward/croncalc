package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func Load() *Config {
	log.Println("Loading environment variables...")
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8010"
		log.Println("PORT not set in environment, using default port 8010")
	} else if _, err := strconv.Atoi(port); err != nil {
		log.Printf("Invalid PORT value: %v", err)
		log.Println("Setting PORT to 8010 as default")
		port = "8010"
	}

	return &Config{
		Port: port,
	}
}
