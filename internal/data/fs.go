package data

import (
	"log"
	"os"
)

// create file when run project
func CreateFile() {
	_, err := os.ReadFile("../../data/data.json")
	if err != nil {
		err = os.WriteFile("../../data/data.json", []byte("{}"), 0644)
		if err != nil {
			log.Fatal(err)
			return
		}
		return
	}
}
