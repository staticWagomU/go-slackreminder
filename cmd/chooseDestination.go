package cmd

import (
	"fmt"
)

func chooseDestination() (string, error) {
	defaultDestination := "me"
	result, err := createPrompt("送信先を入力してください。デフォルトは『"+defaultDestination+"』です。", defaultDestination)
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return result, nil
}
