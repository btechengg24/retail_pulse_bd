package api

import (
	"encoding/json"
	"net/http"
	"RETAIL_PULSE_BD/models"
	"RETAIL_PULSE_BD/services"
)

// SubmitJobHandler handles job submission.
func SubmitJobHandler(w http.ResponseWriter, r *http.Request) {
	// Parse incoming JSON request
	var request models.JobRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Check if count matches the number of visits
	if request.Count != len(request.Visits) {
		http.Error(w, "Mismatch in count and visit length", http.StatusBadRequest)
		return
	}

	// Generate job ID and create job
	jobID := services.CreateJob(request)

	// Respond with the job ID
	response := map[string]interface{}{
		"job_id": jobID,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetJobStatusHandler handles retrieving the status of a job.
func GetJobStatusHandler(w http.ResponseWriter, r *http.Request) {
	jobID := r.URL.Query().Get("jobid")
	if jobID == "" {
		http.Error(w, "Job ID is required", http.StatusBadRequest)
		return
	}

	// Fetch job status
	status, err := services.GetJobStatus(jobID)
	if err != nil {
		http.Error(w, "Job not found", http.StatusBadRequest)
		return
	}

	// Respond with job status
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status)
}
