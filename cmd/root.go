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

	destinations, err := InputDestinations()
	timing, err := chooseTiming()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var remindCommand string

	if timing == OneTime {
		remindCommand, err = BuildOneTimeReminderCommand()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		remindCommand, err = BuildRecurringReminderCommand()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	content, err := inputContent()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	command := fmt.Sprintf("/remind %s \"%s\" %s\n", destinations, content, remindCommand)

	fmt.Println("↓↓↓下の文字列をコピーしてください↓↓↓↓")
	fmt.Println(command)

}
