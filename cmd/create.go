package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Path        string    `json:"path"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Tasks []Task

func TakrProject(name string) (string, error) {
	path := "json/project/" + name
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("mkdir: %w", err)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.WriteFile(path+".json", []byte("[]"), 0644); err != nil {
			log.Fatal(err)
			return "", err
		}
	}

	return path, nil
}

func (t *Tasks) TarkCreate(payload Task) error {
	if payload.Type == "Project" {
		path := "json/project/" + payload.Title

		dataJson, err := os.ReadFile(path)
		if err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("read file: %w", err)
		}

		if len(dataJson) > 0 {
			if err = json.Unmarshal(dataJson, t); err != nil {
				return fmt.Errorf("unmarshal: %w", err)
			}
		}

		if len(*t) > 0 {
			payload.ID = (*t)[len(*t)-1].ID + 1
		} else {
			payload.ID = 1
		}

		now := time.Now()
		payload.Status = "pending"
		payload.CreatedAt = now
		payload.UpdatedAt = now

		*t = append(*t, payload)

		updated, err := json.MarshalIndent(t, "", "  ")
		if err != nil {
			return fmt.Errorf("marshal: %w", err)
		}

		if err = os.WriteFile(path, updated, 0644); err != nil {
			return fmt.Errorf("write file: %w", err)
		}
	}

	return nil
}
