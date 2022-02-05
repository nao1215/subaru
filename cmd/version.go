package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Version is subaru command version.
	Version = "0.1.0"
)

func getVersion() string {
	return fmt.Sprintf("subaru version " + Version + " (under Apache License version 2.0)")
}

var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getVersion())
	},
	Short: "Show subaru command version information",
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
