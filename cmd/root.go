package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "slack-reminder",
	Short: "Slackの`/remind`コマンドをインタラクティブに生成するCLIツール",
	Long:  "Slackの`/remind`コマンドをインタラクティブに生成するCLIツール。",
	Run: func(cmd *cobra.Command, args []string) {
		GenerateReminderCommand()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func GenerateReminderCommand() {
	chosenDestination, err := chooseDestination()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	chosenTiming, err := chooseTiming()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("タイミング：", chosenTiming)

	reminderText, err := reminderText()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("/remind", chosenDestination, chosenTiming, reminderText)
}
