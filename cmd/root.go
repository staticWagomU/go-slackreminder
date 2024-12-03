package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "slack-reminder",
	Short: "Slackの`/remind`コマンドをインタラクティブに生成するCLIツール",
	Long: "Slackの`/remind`コマンドをインタラクティブに生成するCLIツール。",
}



func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

