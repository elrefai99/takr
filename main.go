package main

import (
	"github.com/elrefai99/takr/cmd"
	"github.com/elrefai99/takr/pkg/data"
)

func main() {
	data.CreateFile()
	cmd.TarkReadFiles()
}
