package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type PayloadCreate struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (p *PayloadCreate) Create_Project(payload PayloadCreate) error {
	path := "json/" + payload.Name + ".json"
	fmt.Println(path)
	_, err := os.ReadFile(path)
	if err == nil {
		return err
	}
	if !os.IsNotExist(err) {
		return err
	}

	arr := []Tasks{
		{
			ID:          1,
			Title:       payload.Title,
			Status:      payload.Status,
			Description: payload.Description,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	dataObject, err := json.MarshalIndent(arr, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(path, dataObject, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Success create new task project:", payload.Name)
	return nil
}
