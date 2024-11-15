package models

// Store represents a store from the store master.
type Store struct {
	StoreID   string `json:"store_id"`
	StoreName string `json:"store_name"`
	AreaCode  string `json:"area_code"`
}

// JobRequest represents the request format for submitting jobs.
type JobRequest struct {
	Count  int      `json:"count"`
	Visits []Visit  `json:"visits"`
}

// Visit represents a visit to a store with images.
type Visit struct {
	StoreID   string   `json:"store_id"`
	ImageURLs []string `json:"image_url"`
	VisitTime string   `json:"visit_time"`
}

// JobStatus represents the status of a job.
type JobStatus struct {
	Status  string       `json:"status"`
	JobID   string       `json:"job_id"`
	Error   []JobError   `json:"error,omitempty"`
}

// JobError represents any errors encountered during job processing.
type JobError struct {
	StoreID string `json:"store_id"`
	Error   string `json:"error"`
}
