package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func chooseDestination() (string, error) {
	defaultDestination := "me"
	prompt := promptui.Prompt{
		Label:   "送信先を入力してください。デフォルトは『" + defaultDestination + "』です。",
		Default: defaultDestination,
		Templates: &promptui.PromptTemplates{
			Prompt:  "{{ . }} ",
			Valid:   "{{ . | green }} ",
			Invalid: "{{ . | red }} ",
			Success: "{{ . | bold }} ",
		},
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return result, nil
}
