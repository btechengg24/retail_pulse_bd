package main

import (
	"log"
	"net/http"
	"RETAIL_PULSE_BD/api"
	"RETAIL_PULSE_BD/services"
)

func main() {
	// Load store master data
	if err := services.LoadStoreMaster("store_master.csv"); err != nil {
		log.Fatalf("Failed to load store master: %v", err)
	}

	// Set up HTTP routes
	http.HandleFunc("/api/submit", api.SubmitJobHandler)
	http.HandleFunc("/api/status", api.GetJobStatusHandler)

	// Start the server
	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
