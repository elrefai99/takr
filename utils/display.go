package utils

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/elrefai99/takr/cmd"
)

func PrintResponse(items []cmd.Tasks) {
	format := "%v\t%v\t%v\t%v\t%v\t%v\t\n"

	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "ID", "Title", "Status", "Description", "createdAt", "updatedAt")
	fmt.Fprintf(tw, format, "--", "-----", "------", "-----------", "---------", "---------")

	for _, i := range items {
		desc := i.Description
		if len(desc) > 75 {
			desc = desc[:75] + "..."
		}
		fmt.Fprintf(tw, format, i.ID, i.Title, i.Status, desc, i.CreatedAt, i.UpdatedAt)
	}
	tw.Flush()
}
