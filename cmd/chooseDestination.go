package cmd

import (
	"log"

	"github.com/charmbracelet/huh"
)

func InputDestinations() (string, error) {
	var destination string

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("リマインドする相手は？").
				Description("defaultは me").
				Value(&destination),
		),
	).Run()

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	if destination == "" {
		destination = "me"
	}

	return destination, nil
}
