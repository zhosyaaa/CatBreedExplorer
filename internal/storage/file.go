package storage

import (
	"encoding/json"
	"io/ioutil"
)

func WriteSortedCatBreedsToFile(originGroups map[string][]string) error {
	resultData, err := json.MarshalIndent(originGroups, "", "")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("out.json", resultData, 0644)
	if err != nil {
		return err
	}
	return err
}
