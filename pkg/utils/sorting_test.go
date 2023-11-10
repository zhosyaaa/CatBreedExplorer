package utils

import (
	"CatBreedExplorer/internal/models"
	"reflect"
	"sort"
	"testing"
)

func TestSortCatBreedsByCountryAndLength(t *testing.T) {
	breeds := []models.CatBreed{
		{Country: "USA", Breed: "Maine Coon"},
		{Country: "USA", Breed: "Siamese"},
		{Country: "UK", Breed: "British Shorthair"},
		{Country: "UK", Breed: "Scottish Fold"},
		{Country: "Japan", Breed: "Bobtail"},
	}

	expectedResult := map[string][]string{
		"USA":   {"Siamese", "Maine Coon"},
		"UK":    {"Scottish Fold", "British Shorthair"},
		"Japan": {"Bobtail"},
	}

	result := SortCatBreedsByCountryAndLength(breeds)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Unexpected result. Expected: %v, Got: %v", expectedResult, result)
	}

	for _, names := range result {
		if !sort.SliceIsSorted(names, func(i, j int) bool {
			return len(names[i]) < len(names[j])
		}) {
			t.Errorf("Breeds are not sorted by length for country: %v", names)
		}
	}
}
