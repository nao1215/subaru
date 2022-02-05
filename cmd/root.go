package cmd

import (
	"embed"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// phrases are used to output phrases randomly.
var phrases map[string]string

func init() {
	phrases = make(map[string]string)
}

var rootCmd = &cobra.Command{
	Use:   "subaru",
	Short: "subaru print philosophy or wise sayings.",
	Run: func(cmd *cobra.Command, args []string) {
		os.Exit(randomPrint())
	},
}

func exitError(msg interface{}) {
	fmt.Fprintf(os.Stderr, "subaru: %v\n", msg)
	os.Exit(1)
}

// Execute run subaru command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		exitError(err)
	}
}

// AddCommands add commands from the embedded "*.subaru (text file)"
func AddCommands(fs embed.FS) {
	entries, err := fs.ReadDir("fortune")
	if err != nil {
		exitError(err)
	}

	for _, entry := range entries {
		cmdName := removerExt(entry.Name())
		b, err := fs.ReadFile("fortune/" + entry.Name())
		if err != nil {
			exitError(err)
		}

		rootCmd.AddCommand(&cobra.Command{
			Use:   cmdName,
			Short: "Print phrase related to " + cmdName,
			Run: func(cmd *cobra.Command, args []string) {
				os.Exit(printPhrase(cmd.Name(), string(b)))
			},
		})
		phrases[cmdName] = string(b)
	}
}

func removerExt(path string) string {
	if idx := strings.LastIndexByte(path, '.'); idx >= 0 {
		return path[:idx]
	}
	return path
}

func randomPrint() int {
	for k, v := range phrases {
		return printPhrase(k, v)
	}
	return 1
}

func printPhrase(cmdName, phrase string) int {
	fmt.Println("[" + color.GreenString(cmdName) + "]")
	fmt.Println(phrase)
	return 0
}
