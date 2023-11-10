package app

import (
	"CatBreedExplorer/internal/handlers"
	"CatBreedExplorer/internal/storage"
	"CatBreedExplorer/pkg/utils"
	"github.com/joho/godotenv"
	"log"
)

func Run() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	// Fetch the list of cat breeds
	breeds, err := handlers.GetCatBreeds()
	if err != nil {
		log.Fatal("Error when getting a list of cat breeds:", err)

	}
	// Sort and group breeds by country of origin and name length
	originGroups := utils.SortCatBreedsByCountryAndLength(breeds)

	// Write the sorted data to a file
	err = storage.WriteSortedCatBreedsToFile(originGroups)
	if err != nil {
		log.Fatal("Error when writing data to a file:", err)
	}
	log.Printf("Results successfully written to out.json")
}
