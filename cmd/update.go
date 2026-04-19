package cmd

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type PayloadUpdate struct {
	Title       string
	Status      string
	Description string
}

func UpdateTask(id uint, payload PayloadUpdate) error {
	path := "json/default.json"

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var tasks []Tasks
	if err = json.Unmarshal(data, &tasks); err != nil {
		return err
	}

	found := false
	for i := range tasks {
		if tasks[i].ID != id {
			continue
		}

		if payload.Title != "" {
			tasks[i].Title = payload.Title
		}
		if payload.Status != "" {
			tasks[i].Status = payload.Status
		}
		if payload.Description != "" {
			tasks[i].Description = payload.Description
		}
		tasks[i].UpdatedAt = time.Now()
		found = true
		break
	}

	if !found {
		return errors.New("task not found")
	}

	updated, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, updated, 0644)
}
