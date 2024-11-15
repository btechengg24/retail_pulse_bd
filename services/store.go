package services

import (
	"encoding/csv"
	"os"
	"RETAIL_PULSE_BD/models"
)

var storeMaster map[string]models.Store

// LoadStoreMaster loads the store master data from a CSV file.
func LoadStoreMaster(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	storeMaster = make(map[string]models.Store)
	for _, record := range records {
		store := models.Store{
			StoreID:   record[0],
			StoreName: record[1],
			AreaCode:  record[2],
		}
		storeMaster[store.StoreID] = store
	}

	return nil
}
