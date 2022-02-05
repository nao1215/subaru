package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	// Version is subaru command version.
	Version = "0.0.1"
)

func getVersion() string {
	return fmt.Sprintf("Version: %s\nOS     : %s\nArch   : %s",
		Version, runtime.GOOS, runtime.GOARCH)
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
