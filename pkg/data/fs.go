package data

import (
	"log"
	"os"
	"path/filepath"
)

// create file when run project
func CreateFile() {
	dir := filepath.Dir("json/data.json")
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatal(err)
		return
	}

	_, err := os.ReadFile("json/data.json")
	if err != nil {
		err = os.WriteFile("json/data.json", []byte("{}"), 0644)
		if err != nil {
			log.Fatal(err)
			return
		}
		return
	}
}
