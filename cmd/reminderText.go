package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func reminderText() (string, error) {
	// 未入力はinvalid
	prompt := promptui.Prompt{
		Label: "リマインドする内容を入力してください",
		Validate: func(s string) error {
			if len(s) < 1 {
				return fmt.Errorf("1文字以上入力してください")
			}
			return nil
		},
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return result, nil
}
