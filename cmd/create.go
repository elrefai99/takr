package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Tasks struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type PayloadCreate struct {
	Title       string
	Status      string
	Description string
}

func (p *PayloadCreate) Create_Project(payload PayloadCreate) error {
	path := "json/default.json"

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var tasks []Tasks
	if err = json.Unmarshal(data, &tasks); err != nil {
		return err
	}

	var nextID uint = 1
	if len(tasks) > 0 {
		nextID = tasks[len(tasks)-1].ID + 1
	}

	tasks = append(tasks, Tasks{
		ID:          nextID,
		Title:       payload.Title,
		Status:      payload.Status,
		Description: payload.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})

	dataObject, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	if err = os.WriteFile(path, dataObject, 0644); err != nil {
		return err
	}

	fmt.Printf("Task created with ID: %d\n", nextID)
	return nil
}
