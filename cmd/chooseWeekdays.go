package cmd

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
)

func ChooseMultiWeekdays() (string, error) {
	var command []string

	options := []huh.Option[string]{
		huh.NewOption("日曜日", "Sunday"),
		huh.NewOption("月曜日", "Monday"),
		huh.NewOption("火曜日", "Tuesday"),
		huh.NewOption("水曜日", "Wednesday"),
		huh.NewOption("木曜日", "Thursday"),
		huh.NewOption("金曜日", "Friday"),
		huh.NewOption("土曜日", "Saturday"),
	}

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("リマインドする曜日は？").
				Options(options...).
				Description("複数選択可能").
				Validate(func(s []string) error {
					if len(s) == 0 {
						return fmt.Errorf("少なくとも1つの曜日を選択してください")
					}
					return nil
				}).
				Value(&command),
		),
	).Run()

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	// 週の曜日を指定するためには、曜日の英語名をカンマ区切りで指定する
	var weekdays = ""
	for _, c := range command {
		weekdays += c + ", "
	}
	if len(weekdays) > 0 {
		// 末尾のカンマとスペースを削除
		weekdays = weekdays[:len(weekdays)-2]
	}

	if len(command) == 7 {
		weekdays = "every day"
	}

	return weekdays, nil
}

func ChooseWeekday() (string, error) {
	var command string

	options := []huh.Option[string]{
		huh.NewOption("日曜日", "Sunday"),
		huh.NewOption("月曜日", "Monday"),
		huh.NewOption("火曜日", "Tuesday"),
		huh.NewOption("水曜日", "Wednesday"),
		huh.NewOption("木曜日", "Thursday"),
		huh.NewOption("金曜日", "Friday"),
		huh.NewOption("土曜日", "Saturday"),
	}

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("リマインドする曜日は？").
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
