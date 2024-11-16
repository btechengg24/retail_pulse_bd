package services

import (
	"errors"
	"log"
	"RETAIL_PULSE_BD/models"
	"sync"
)

var (
	jobs      = sync.Map{}
	jobIDLock sync.Mutex
	jobIDSeq  int
)

// CreateJob creates a new job and processes it asynchronously
func CreateJob(request models.JobRequest) int {
	// Generate a unique job ID
	jobIDLock.Lock()
	jobIDSeq++
	jobID := jobIDSeq
	jobIDLock.Unlock()

	// Store job initially as "processing"
	jobs.Store(jobID, "processing")

	// Process the job asynchronously
	go processJob(jobID, request)

	return jobID
}

// GetJobStatus retrieves the status of a job
func GetJobStatus(jobID int) (string, error) {
	status, exists := jobs.Load(jobID)
	if !exists {
		return "", errors.New("job not found")
	}
	return status.(string), nil
}

func processJob(jobID int, request models.JobRequest) {
	for _, visit := range request.Visits {
		log.Printf("Processing visit for store %s with %d images", visit.StoreID, len(visit.ImageURLs))
		processImages(visit.ImageURLs)
	}

	// Mark the job as completed
	jobs.Store(jobID, "completed")
	log.Printf("Job %d completed", jobID)
}

func processImages(imageURLs []string) {
	for _, imageURL := range imageURLs {
		log.Printf("Processing image: %s", imageURL)
		// Simulate processing
	}
}
