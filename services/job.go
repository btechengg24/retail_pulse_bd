package services

import (
	"errors"
	"sync"
	"time"

	"RETAIL_PULSE_BD/models"
)

// Job represents a processing job.
type Job struct {
	Status string
	Error  []models.JobError
}

var (
	jobs      = make(map[string]Job)
	jobsMutex sync.Mutex
)

// CreateJob creates and starts a new job.
func CreateJob(request models.JobRequest) string {
	jobID := generateJobID()

	jobsMutex.Lock()
	jobs[jobID] = Job{
		Status: "Processing",
	}
	jobsMutex.Unlock()

	go processJob(jobID, request)

	return jobID
}

// GetJobStatus retrieves the status of a job.
func GetJobStatus(jobID string) (models.JobStatus, error) {
	jobsMutex.Lock()
	defer jobsMutex.Unlock()

	job, exists := jobs[jobID]
	if !exists {
		return models.JobStatus{}, errors.New("job not found")
	}

	return models.JobStatus{
		JobID:  jobID,
		Status: job.Status,
		Error:  job.Error,
	}, nil
}

func processJob(jobID string, request models.JobRequest) {
	time.Sleep(2 * time.Second) // Simulate processing delay

	var errors []models.JobError

	for _, visit := range request.Visits {
		if visit.StoreID == "" {
			errors = append(errors, models.JobError{
				StoreID: visit.StoreID,
				Error:   "Store ID is missing",
			})
		}
	}

	jobsMutex.Lock()
	job := jobs[jobID]
	job.Status = "Completed"
	job.Error = errors
	jobs[jobID] = job
	jobsMutex.Unlock()
}

func generateJobID() string {
	return "JOB_" + time.Now().Format("20060102150405")
}
