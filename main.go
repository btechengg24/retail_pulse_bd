package main

import (
	"log"
	"net/http"
	"RETAIL_PULSE_BD/api"
)

func main() {
	// Map routes to handlers
	http.HandleFunc("/api/submit", api.SubmitJobHandler)
	http.HandleFunc("/api/submit/", api.SubmitJobHandler)
	http.HandleFunc("/api/status", api.GetJobStatusHandler)

	// Start the HTTP server
	log.Println("Server is running on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
