package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Data struct {
	ID int `json:"id"`
}

func TarkReadFiles() {
	data, err := os.ReadFile("json/data.json")
	if err != nil {
		log.Fatal(err)
		return
	}
	var item []Data
	err = json.Unmarshal(data, &item)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%+v\n", item)
}
