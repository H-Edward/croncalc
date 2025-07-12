package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/H-Edward/croncalc/config"
	"github.com/H-Edward/croncalc/handlers"
	"github.com/H-Edward/croncalc/services"
)

func main() {
	cfg := config.Load()

	// Initialize timezones at startup
	_, err := services.InitializeTimezones()
	if err != nil {
		log.Printf("Warning: Failed to initialize timezones: %v", err)
	}

	// Prepare timezone response cache at startup
	handlers.PrepareTimezonesResponse()

	http.HandleFunc("/api/parse", handlers.ParseHandler)
	http.HandleFunc("/api/timezones", handlers.AvailableTimezonesHandler)

	// Serve static from static/
	http.Handle("/", http.FileServer(http.Dir("static/")))

	serverAddr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Starting server on port %s", cfg.Port)
	if err := http.ListenAndServe(serverAddr, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
