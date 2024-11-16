package api

import (
	"encoding/json"
	"net/http"
	"RETAIL_PULSE_BD/models"
	"RETAIL_PULSE_BD/services"
	"strconv"
)

// SubmitJobHandler handles job submission
func SubmitJobHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	var jobRequest models.JobRequest
	err := json.NewDecoder(r.Body).Decode(&jobRequest)
	if err != nil || jobRequest.Count != len(jobRequest.Visits) {
		http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
		return
	}

	// Generate and process the job
	jobID := services.CreateJob(jobRequest)

	// Return the job ID in the response
	response := map[string]interface{}{"job_id": jobID}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetJobStatusHandler retrieves the status of a job
func GetJobStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	// Extract job ID from query parameters
	jobIDStr := r.URL.Query().Get("jobid")
	jobID, err := strconv.Atoi(jobIDStr)
	if err != nil {
		http.Error(w, `{"error": "Invalid job ID"}`, http.StatusBadRequest)
		return
	}

	// Fetch the job status
	status, err := services.GetJobStatus(jobID)
	if err != nil {
		http.Error(w, `{"error": "Job ID not found"}`, http.StatusBadRequest)
		return
	}

	// Return the job status
	response := map[string]interface{}{
		"job_id": jobID,
		"status": status,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
