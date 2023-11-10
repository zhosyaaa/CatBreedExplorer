package utils

import (
	"CatBreedExplorer/internal/models"
	"sort"
)

func SortCatBreedsByCountryAndLength(breeds []models.CatBreed) map[string][]string {
	originGroups := make(map[string][]string)
	for _, breed := range breeds {
		originGroups[breed.Country] = append(originGroups[breed.Country], breed.Breed)
	}
	for _, names := range originGroups {
		sort.Slice(names, func(i, j int) bool {
			return len(names[i]) < len(names[j])
		})
	}
	return originGroups
}
