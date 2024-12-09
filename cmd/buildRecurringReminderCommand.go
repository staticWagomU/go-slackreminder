package cmd

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
)

func BuildRecurringReminderCommand() (string, error) {
	var command string

	const (
		// 毎日
		EveryDay = "every day"
		// 平日
		EveryWeekday = "every weekday"
		// 毎週○曜日
		Weekdays = "weekdays"
		// 隔週の○曜日
		AlternateWeekdays = "alternate weekdays"
		// 毎月○日
		EveryMonth = "every month"
		// 隔月の○日
		AlternateMonths = "alternate months"
		// 毎年○月○日
		EveryYear = "every year"
	)

	options := []huh.Option[string]{
		huh.NewOption("毎日", EveryDay),
		huh.NewOption("平日", EveryWeekday),
		huh.NewOption("毎週○曜日", Weekdays),
		huh.NewOption("隔週の○曜日", AlternateWeekdays),
		huh.NewOption("毎月○日", EveryMonth),
		huh.NewOption("隔月の○日", AlternateMonths),
		huh.NewOption("毎年○月○日", EveryYear),
	}

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("リマインドする日時は？").
				Options(options...).
				Value(&command),
		),
	).Run()

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	if command == EveryDay || command == EveryWeekday {
		return fmt.Sprintf("%s", command), nil
	}

	switch command {
	case Weekdays:
		command, err = ChooseMultiWeekdays()
		if err != nil {
			log.Fatal(err)
			return "", err
		}
		command = fmt.Sprintf("every %s", command)
	case AlternateWeekdays:
		command, err = ChooseWeekday()
		if err != nil {
			log.Fatal(err)
			return "", err
		}
		command = fmt.Sprintf("every other %s", command)
	case EveryMonth:
		command, err = InputDay()
		if err != nil {
			log.Fatal(err)
			return "", err
		}
		command = fmt.Sprintf("on the %s of every month", command)
	case AlternateMonths:
		command, err = InputDay()
		if err != nil {
			log.Fatal(err)
			return "", err
		}
		command = fmt.Sprintf("every other %s", command)
	case EveryYear:
		command, err = ChooseMonth()
		if err != nil {
			log.Fatal(err)
			return "", err
		}
		day, err := InputDay()
		if err != nil {
			log.Fatal(err)
			return "", err
		}

		command = fmt.Sprintf("every %s %s", command, day)
	default:
		return "", fmt.Errorf("invalid command: %s", command)
	}

	time, err := InputAbsoluteTime()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	command = fmt.Sprintf("%s %s", time, command)

	return command, nil
}
