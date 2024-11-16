package models

// Store represents a store's details
type Store struct {
	StoreID  string
	Name     string
	AreaCode string
}

// JobRequest represents the structure of the incoming job submission
type JobRequest struct {
	Count  int `json:"count"`
	Visits []struct {
		StoreID   string   `json:"store_id"`
		ImageURLs []string `json:"image_url"`
		VisitTime string   `json:"visit_time"`
	} `json:"visits"`
}
