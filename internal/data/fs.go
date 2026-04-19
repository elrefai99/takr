package create_default

import (
	"encoding/json"
	"fmt"
	"log"
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

func Create_default() {
	_, err := os.ReadFile("json/default.json")
	if err != nil {
		arr := []Tasks{}
		data := Tasks{}

		arr = append(arr, data)
		dataObject, err := json.Marshal(arr)
		if err != nil {
			log.Fatal(err)
			return
		}
		os.WriteFile("json/default.json", dataObject, 0644)
		return
	}
	fmt.Println("Wellcome in Takr CLI ...>")
}
