package main

import (
	"embed"

	"github.com/nao1215/subaru/cmd"
)

//go:embed fortune/*.subaru
var fortune embed.FS

func main() {
	cmd.AddCommands(fortune)
	cmd.Execute()
}
