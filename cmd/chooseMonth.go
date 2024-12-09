package cmd

import (
	"log"

	"github.com/charmbracelet/huh"
)

func ChooseMonth() (string, error) {
	var command string

	options := []huh.Option[string]{
		huh.NewOption("1月", "January"),
		huh.NewOption("2月", "February"),
		huh.NewOption("3月", "March"),
		huh.NewOption("4月", "April"),
		huh.NewOption("5月", "May"),
		huh.NewOption("6月", "June"),
		huh.NewOption("7月", "July"),
		huh.NewOption("8月", "August"),
		huh.NewOption("9月", "September"),
		huh.NewOption("10月", "October"),
		huh.NewOption("11月", "November"),
		huh.NewOption("12月", "December"),
	}

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("リマインドする月は？").
				Options(options...).
				Value(&command),
		),
	).Run()

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return command, nil
}
