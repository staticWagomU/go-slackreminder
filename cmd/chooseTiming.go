package cmd

import (
	"log"

	"github.com/charmbracelet/huh"
)

const (
	// OneTime 一回限り
	OneTime = "onetime"
	// Recurring 定例
	Recurring = "recurring"
)

func chooseTiming() (string, error) {
	var timing string
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("リマインド回数は？").
				Options(
					huh.NewOption("一回限り", OneTime),
					huh.NewOption("定例", Recurring),
				).
				Value(&timing),
		),
	).Run()

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return timing, nil
}
