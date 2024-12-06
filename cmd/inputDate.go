package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/charmbracelet/huh"
)

func InputDate() (string, error) {
	var date string

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("送信日は？\nyyyy/MM/ddかMM/ddで書いてね").
				Validate(func(s string) error {
					if s == "" {
						return nil
					}
					_, err1 := time.Parse("2006/01/02", s)
					_, err2 := time.Parse("01/02", s)
					if err1 != nil && err2 != nil {
						return fmt.Errorf("正しい形式で入力してください")
					}

					return nil
				}).
				Value(&date),
		),
	).Run()

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	_, err = time.Parse("01/02", date)

	if err != nil {
		date = fmt.Sprintf("%s/%s", time.Now().Format("2006"), date)
	}

	return date, nil
}
