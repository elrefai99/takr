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

func main() {
	create_default.Create_default()
	flag.Parse()

	args := flag.Args()

	switch args[0] {
	case "create":
		title := readLine("Please input title of Task ...> ")
		status := readLine("Please choose type of status (To-Do, In Progress, Done) ...> ")
		description := readLine("Please input description of Task ...> ")

		p := cmd.PayloadCreate{
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
		data, err := cmd.GetData()
		if err != nil {
			log.Fatal(err)
			return
		}
		utils.PrintResponse(data)
		return
	case "update":
		var idInput uint
		fmt.Print("Please input Task ID to update ...> ")
		fmt.Scanln(&idInput)

		title := readLine("New title (leave empty to skip) ...> ")
		status := readLine("New status (leave empty to skip) ...> ")
		description := readLine("New description (leave empty to skip) ...> ")

		err := cmd.UpdateTask(idInput, cmd.PayloadUpdate{
			Title:       title,
			Status:      status,
			Description: description,
		})
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("Task updated successfully.")
		return
	default:
		fmt.Println("unknown command:", args[0])
		return
	}
}
