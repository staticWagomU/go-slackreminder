package cmd

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
)

func BuildOneTimeReminderCommand() (string, error) {
	var command string

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("リマインドする日時は？").
				Options(
					huh.NewOption("20分後", "20 minutes"),
					huh.NewOption("1時間後", "1 hour"),
					huh.NewOption("3時間後", "3 hours"),
					huh.NewOption("明日", "1 day"),
					huh.NewOption("来週", "1 week"),
					huh.NewOption("来月", "1 month"),
					huh.NewOption("特定の日付？", "absolute"),
					huh.NewOption("○[分時ヶ月]後?", "relative"),
				).
				Value(&command),
		),
	).Run()

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	if command != "absolute" && command != "relative" {
		return fmt.Sprintf("in %s", command), nil
	}

	var date = ""
	var time = ""
	if command == "absolute" {
		time, err = InputRelativeTime()
		if err != nil {
			log.Fatal(err)
			return "", err
		}
	} else {
		date, err = InputDate()
		if err != nil {
			log.Fatal(err)
			return "", err
		}

		time, err = InputAbsoluteTime()
		if err != nil {
			log.Fatal(err)
			return "", err
		}
	}

	command = ""

	if date != "" {
		command = fmt.Sprintf("on %s", date)
	}

	if time != "" {
		command = fmt.Sprintf("%s in %s", command, time)
	}

	return command, nil
}
