package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/elrefai99/takr/cmd"
	create_default "github.com/elrefai99/takr/internal/data"
	"github.com/elrefai99/takr/utils"
)

var scanner = bufio.NewScanner(os.Stdin)

func readLine(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func create(title string, status string, description string) error {
	p := cmd.PayloadCreate{
		Title:       title,
		Status:      status,
		Description: description,
	}

	err := p.Create_Project(p)
	if err != nil {
		return err
	}
	return nil
}

func get() {
	data, err := cmd.GetData()
	if err != nil {
		log.Fatal(err)
		return
	}
	utils.PrintResponse(data)
}

func update(title string, status string, description string, idInput uint) error {
	err := cmd.UpdateTask(idInput, cmd.PayloadUpdate{
		Title:       title,
		Status:      status,
		Description: description,
	})
	if err != nil {
		return err
	}
	fmt.Println("Task updated successfully.")
	return nil
}

func main() {
	create_default.Create_default()
	flag.Parse()

	args := flag.Args()

	switch args[0] {
	case "create":
		title := readLine("Please input title of Task ...> ")
		status := readLine("Please choose type of status (To-Do, In Progress, Done) ...> ")
		description := readLine("Please input description of Task ...> ")
		err := create(title, status, description)
		if err != nil {
			log.Fatal(err)
			return
		}
		return
	case "get":
		get()
		return
	case "update":
		var idInput uint
		fmt.Print("Please input Task ID to update ...> ")
		fmt.Scanln(&idInput)
		title := readLine("Please input title of Task ...> ")
		status := readLine("Please choose type of status (To-Do, In Progress, Done) ...> ")
		description := readLine("Please input description of Task ...> ")
		err := update(title, status, description, idInput)
		if err != nil {
			log.Fatal(err)
			return
		}
		return

	case "delete":
		var id uint
		fmt.Print("Please input Task ID to update ...> ")
		fmt.Scanln(&id)
		err := cmd.DeleteTask(uint(id))
		if err != nil {
			log.Fatal(err)
			return
		}
		return
	default:
		fmt.Println("unknown command:", args[0])
		return
	}
}
