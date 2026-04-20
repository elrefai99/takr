package cmd

import (
	"encoding/json"
	"fmt"
	"os"
)

func DeleteTask(id uint) error {
	path := "json/default.json"

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var tasks []Tasks
	if err = json.Unmarshal(data, &tasks); err != nil {
		return err
	}

	filtered := make([]Tasks, 0, len(tasks)-int(id))
	found := false

	for _, t := range tasks {
		if t.ID == id {
			found = true
			continue
		}
		filtered = append(filtered, t)
	}

	if !found {
		return fmt.Errorf("task with id %d not found in project: ", id)
	}

	out, err := json.MarshalIndent(filtered, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, out, 0644)
}
