package services

import (
	"errors"
	"RETAIL_PULSE_BD/models"
	"sync"
)

var storeMaster = map[string]models.Store{
	"S00339218": {"S00339218", "Store A", "12345"},
	"S01408764": {"S01408764", "Store B", "67890"},
}

var storeMutex sync.RWMutex

// GetStore retrieves a store by ID
func GetStore(storeID string) (models.Store, error) {
	storeMutex.RLock()
	defer storeMutex.RUnlock()

	store, exists := storeMaster[storeID]
	if !exists {
		return models.Store{}, errors.New("store not found")
	}
	return store, nil
}
