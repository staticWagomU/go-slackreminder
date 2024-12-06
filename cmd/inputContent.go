package cmd

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
)

func inputContent() (string, error) {
	var content string

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewText().
				Title("リマインドする内容は？").
				Validate(func(s string) error {
					if len(s) == 0 {
						return fmt.Errorf("内容を入力してください")
					}
					return nil
				}).
				Value(&content),
		),
	).Run()

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return content, nil
}
