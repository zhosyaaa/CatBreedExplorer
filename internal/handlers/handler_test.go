package handlers

import (
	"CatBreedExplorer/internal/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetCatBreeds(t *testing.T) {
	apiResponse := map[string]interface{}{
		"data": []interface{}{
			map[string]interface{}{
				"breed":   "Siamese",
				"country": "Thailand",
				"origin":  "Natural",
				"coat":    "Short",
				"pattern": "Colorpoint",
			},
			map[string]interface{}{
				"breed":   "Maine Coon",
				"country": "USA",
				"origin":  "Natural",
				"coat":    "Long",
				"pattern": "Tabby",
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(apiResponse)
	}))
	defer server.Close()

	os.Setenv("API_URL", server.URL)
	defer os.Unsetenv("API_URL")

	breeds, err := GetCatBreeds()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedBreeds := []models.CatBreed{
		{
			Breed:   "Siamese",
			Country: "Thailand",
			Origin:  "Natural",
			Coat:    "Short",
			Pattern: "Colorpoint",
		},
		{
			Breed:   "Maine Coon",
			Country: "USA",
			Origin:  "Natural",
			Coat:    "Long",
			Pattern: "Tabby",
		},
	}

	if len(breeds) != len(expectedBreeds) {
		t.Errorf("Unexpected result length. Expected: %d, Got: %d", len(expectedBreeds), len(breeds))
	}

	for i, expected := range expectedBreeds {
		if expected != breeds[i] {
			t.Errorf("Unexpected breed at index %d. Expected: %v, Got: %v", i, expected, breeds[i])
		}
	}
}
