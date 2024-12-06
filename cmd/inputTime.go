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

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("リマインドする日時は？").
				Description("使える単位は \n- minutes\n- hours\n- days\n- weeks\n例：10 minutes").
				Validate(func(s string) error {
					if s == "" {
						return nil
					}

					relativeUnits := []string{"minutes", "hours", "days", "weeks"}
					parts := strings.Split(s, " ")
					if len(parts) != 2 {
						return fmt.Errorf("正しい形式で入力してください")
					}
					_, err := strconv.Atoi(parts[0])
					if err != nil {
						return err
					}

					validUnit := false
					for _, unit := range relativeUnits {
						if parts[1] == unit {
							validUnit = true
							break
						}
					}

					if !validUnit {
						return fmt.Errorf("正しい単位を入力してください")
					}

					return nil
				}).
				Value(&command),
		),
	)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return command, nil
}

func InputAbsoluteTime() (string, error) {
	var command string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("リマインドする時間は？").
				Description("24時間のhh:mm形式で書いてね").
				Validate(func(s string) error {
					if s == "" {
						return nil
					}

					_, err := time.Parse("15:04", s)
					return err
				}).
				Value(&command),
		),
	)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return command, nil
}
