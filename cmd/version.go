package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of slack-reminder",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version 1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
