package main

import (
	"flag"
	"fmt"

	"github.com/elrefai99/takr/cmd"
	"github.com/elrefai99/takr/pkg/data"
)

func main() {

	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		data.CreateFile()
		return
	}

	// folders
	switch args[0] {
	case "login":
		fmt.Println("login flow")
	case "create":
		fmt.Println("create flow >")

		// create project ...
		fmt.Println("Do you want to create a new project? (y/n)")
		var choice string
		fmt.Scanln(&choice)

		var name string
		if choice == "y" {
			fmt.Println("Add project name:")
			fmt.Scanln(&name)
			path, err := cmd.TakrProject(name)
			if err != nil {
				fmt.Println("Error creating project:", err)
				return
			}
			fmt.Println("Project created at:", path)
		}

		// create database ...
		fmt.Println("Do you want to create a new database? (y/n)")
		var database string
		fmt.Scanln(&database)
		if database == "y" {
			TarkCreate := cmd.Tasks{}
			var title string
			var description string
			var status string
			fmt.Println("Add task title:")
			fmt.Scanln(&title)
			fmt.Println("Add task description:")
			fmt.Scanln(&description)
			fmt.Println("Add task status (To Do, In Progress, Done):")
			fmt.Scanln(&status)
			err := TarkCreate.TarkCreate(cmd.Task{
				Path:        name,
				Title:       title,
				Status:      status,
				Description: description,
				Type:        "Project",
			})
			if err != nil {
				fmt.Println("Error creating task:", err)
				return
			}
			fmt.Println("Database created at:")
		}
		// ...
	case "update":
		cmd.TarkReadFiles()

	case "read":
		cmd.TarkReadFiles()
	default:
		fmt.Println("unknown command:", args[0])
	}
}
