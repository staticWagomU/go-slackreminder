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
				Validate(validateDate).
				Value(&date),
		),
	).Run()

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return date, nil
}

func validateDate(s string) error {
	if _, err1 := time.Parse("2006/01/02", s); err1 == nil {
		return nil
	}
	if _, err2 := time.Parse("01/02", s); err2 == nil {
		return nil
	}
	return fmt.Errorf("正しい形式で入力してください")
}

func InputDay() (string, error) {
	var day string

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("送信日は？\nyyyy/MM/ddかMM/ddで書いてね").
				Validate(validateDay).
				Value(&day),
		),
	).Run()

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	if day == "" {
		day = time.Now().Format("2006/01/02")
	}

	return fmt.Sprintf("%sth", day), nil
}

func validateDay(s string) error {
	if s == "" {
		return nil
	}
	if _, err := time.Parse("02", s); err != nil {
		return fmt.Errorf("正しい形式で入力してください")
	}
	return nil
}
