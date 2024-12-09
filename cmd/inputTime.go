package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/huh"
)

func InputRelativeTime() (string, error) {
	var command string

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("リマインドする日時は？").
				Description("使える単位は \n- minutes\n- hours\n- days\n- weeks\n例：10 minutes").
				Validate(validateRelativeTime).
				Value(&command),
		),
	).Run()

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return command, nil
}

func validateRelativeTime(s string) error {
	if s == "" {
		return nil
	}

	relativeUnits := map[string]bool{"minutes": true, "hours": true, "days": true, "weeks": true}
	parts := strings.Split(s, " ")
	if len(parts) != 2 {
		return fmt.Errorf("正しい形式で入力してください")
	}
	if _, err := strconv.Atoi(parts[0]); err != nil {
		return err
	}
	if !relativeUnits[parts[1]] {
		return fmt.Errorf("正しい単位を入力してください")
	}

	return nil
}

func InputAbsoluteTime() (string, error) {
	var command string

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("リマインドする時間は？").
				Description("24時間のhh:mm形式で書いてね").
				Validate(validateAbsoluteTime).
				Value(&command),
		),
	).Run()

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return command, nil
}

func validateAbsoluteTime(s string) error {
	if s == "" {
		return nil
	}

	_, err := time.Parse("15:04", s)
	return err
}
