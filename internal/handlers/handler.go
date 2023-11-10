package handlers

import (
	"CatBreedExplorer/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func GetCatBreeds() ([]models.CatBreed, error) {
	apiURL := os.Getenv("API_URL")
	response, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("incorrect response status: %s", response.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}

	var breeds []models.CatBreed
	if data, ok := result["data"].([]interface{}); ok {
		for _, item := range data {
			if catBreed, ok := item.(map[string]interface{}); ok {
				breed := models.CatBreed{
					Breed:   catBreed["breed"].(string),
					Country: catBreed["country"].(string),
					Origin:  catBreed["origin"].(string),
					Coat:    catBreed["coat"].(string),
					Pattern: catBreed["pattern"].(string),
				}
				breeds = append(breeds, breed)
			}
		}
	}

	return breeds, nil
}
