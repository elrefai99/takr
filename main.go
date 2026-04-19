package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/elrefai99/takr/cmd"
	"github.com/elrefai99/takr/utils"
)

func main() {
	cmd.Create_default()

	flag.Parse()

	args := flag.Args()

	switch args[0] {
	case "create":
		var name string
		var title string
		var status string
		var description string

		fmt.Println("Please input Name of Task ...>")
		fmt.Scanln(&name)
		fmt.Println("Please input title of Task ...>")
		fmt.Scanln(&title)
		fmt.Println("Please choose type of status (To-Do, In Progress, Done) ...>")
		fmt.Scanln(&status)
		fmt.Println("Please input description of Task ...>")
		fmt.Scanln(&description)
		p := cmd.PayloadCreate{
			Name:        name,
			Title:       title,
			Status:      status,
			Description: description,
		}

		err := p.Create_Project(p)
		if err != nil {
			log.Fatal(err)
			return
		}
		return
	case "get":
		var name string
		fmt.Println("Please input Name of Task ...>")
		fmt.Scanln(&name)
		data, err := cmd.GetData(name)
		if err != nil {
			log.Fatal(err)
			return
		}
		utils.PrintResponse(data)
		return
	default:
		fmt.Println("unknown command:", args[0])
		return
	}
}
