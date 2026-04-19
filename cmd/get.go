package cmd

import (
	"encoding/json"
	"os"
)

func GetData(name string) ([]Tasks, error) {
	path := "json/" + name + ".json"

	taskData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var item []Tasks

	err = json.Unmarshal(taskData, &item)
	if err != nil {
		return nil, err
	}

	return item, nil
}
