package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type ParseResponse struct {
	Expr  string   `json:"expr"`
	Next5 []string `json:"next5"`
}

type AvailableTimezonesResponse struct {
	Timezones []Territory `json:"timezones"`
}
type Territory struct {
	Name    string   `json:"name"`
	Regions []string `json:"regions"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}

func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	RespondWithJSON(w, statusCode, ErrorResponse{Error: message})
	log.Printf("Error response (%d): %s", statusCode, message)
}
