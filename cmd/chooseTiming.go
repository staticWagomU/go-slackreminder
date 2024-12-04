package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
)

type ReminderType struct {
	Label string
	Value string
}

type TimeOption struct {
	Label string
	Value string
}

type item struct {
	ID         string
	IsSelected bool
}

const (
	RelativeReminder  = "relative"
	AbsoluteReminder  = "absolute"
	RecurringReminder = "recurring"
)

func chooseTiming() (string, error) {
	items := []*item{
		{
			ID: "Cat",
		},
		{
			ID: "Dog",
		},
		{
			ID: "Horse",
		},
		{
			ID: "Parrot",
		},
		{
			ID: "Zebra",
		},
	}
	selected, err := selectItems(0, items)
	if err != nil {
		log.Fatal(err)
	}

	// Print selected items
	fmt.Println("Selected:")
	for _, s := range selected {
		fmt.Println(s.ID)
	}

	reminderType := []ReminderType{
		{Label: "相対時間", Value: AbsoluteReminder},
		{Label: "絶対時間", Value: RelativeReminder},
		{Label: "定期的リマインド", Value: RecurringReminder},
	}

	i, err := createSelectPrompt("リマインダーの種類を選択", reminderType)
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	var reminderTypeResult string
	switch reminderType[i].Value {
	case RelativeReminder:
		reminderTypeResult = handleRelativeReminder()
	case AbsoluteReminder:
		reminderTypeResult = handleAbsoluteReminder()
	case RecurringReminder:
		reminderTypeResult = handleRecurringReminder()
	default:
		log.Fatalf("不明なリマインダータイプ: %s", reminderType[i].Value)
	}

	return reminderTypeResult, nil
}


func handleRelativeReminder() string {
	relativeOptions := []TimeOption{
		{Label: "30秒後", Value: "in 30 seconds"},
		{Label: "30分後", Value: "in 30 minutes"},
		{Label: "1時間後", Value: "in 1 hour"},
		{Label: "カスタム時間入力", Value: "custom"},
	}

	i, err := createSelectPrompt("相対時間を選択", relativeOptions)
	if err != nil {
		log.Fatalf("プロンプト選択エラー: %v", err)
	}

	relativeUnits := []string{"seconds", "minutes", "hours"}
	if relativeOptions[i].Value == "custom" {
		customPrompt := promptui.Prompt{
			Label: "時間と単位を入力 (数字 単位, 単位は右記参照 minutes, seconds, hours)",
			Validate: func(input string) error {
				parts := strings.Split(input, " ")
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
			},
		}

		customTime, err := customPrompt.Run()
		if err != nil {
			log.Fatalf("カスタム時間入力エラー: %v", err)
		}
		return fmt.Sprintf("in %s", customTime)
	}

	return relativeOptions[i].Value
}


func handleAbsoluteReminder() string {
	absoluteOptions := []TimeOption{
		{Label: "日付", Value: "date"},
		{Label: "時刻", Value: "time"},
		{Label: "日付と時刻", Value: "datetime"},
	}

	i, err := createSelectPrompt("種類を選択", absoluteOptions)
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	switch absoluteOptions[i].Value {
	case "date":
		datePrompt := promptui.Prompt{
			Label:   "日付を入力してください (yyyy/mm/dd)",
			Default: time.Now().Format("2006/01"),
			Validate: func(input string) error {
				_, err := time.Parse("2006/01/02", input)
				return err
			},
		}

		dateResult, err := datePrompt.Run()
		if err != nil {
			log.Fatalf("日付入力エラー: %v", err)
		}
		return fmt.Sprintf("on %s", dateResult)
	case "time":
		timePrompt := promptui.Prompt{
			Label: "時刻を入力してください (HH:MM 24時間形式)",
			Validate: func(input string) error {
				_, err := time.Parse("15:04", input)
				return err
			},
		}

		timeResult, err := timePrompt.Run()
		if err != nil {
			log.Fatalf("時刻入力エラー: %v", err)
		}
		return fmt.Sprintf("at %s", timeResult)
	case "datetime":
		datePrompt := promptui.Prompt{
			Label:   "日付を入力してください (yyyy/mm/dd)",
			Default: time.Now().Format("2006/01"),
			Validate: func(input string) error {
				_, err := time.Parse("2006/01/02", input)
				return err
			},
		}

		dateResult, err := datePrompt.Run()
		if err != nil {
			log.Fatalf("日付入力エラー: %v", err)
		}

		timePrompt := promptui.Prompt{
			Label: "時刻を入力してください (HH:MM 24時間形式)",
			Validate: func(input string) error {
				_, err := time.Parse("15:04", input)
				return err
			},
		}

		timeResult, err := timePrompt.Run()
		if err != nil {
			log.Fatalf("時刻入力エラー: %v", err)
		}
		return fmt.Sprintf("at %s on %s", timeResult, dateResult)
	}

	return ""
}

func handleRecurringReminder() string {
	recurringOptions := []TimeOption{
		{Label: "毎日", Value: "every day"},
		{Label: "平日", Value: "weekdays"},
		{Label: "毎週", Value: "every week"},
		{Label: "毎月", Value: "every month"},
		{Label: "毎年", Value: "every year"},
		{Label: "曜日指定(複数も可能)", Value: "advanced"},
		{Label: "カスタム間隔", Value: "interval"},
	}

	i, err := createSelectPrompt("リマインダーの種類を選択", recurringOptions)
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	if recurringOptions[i].Value == "advanced" {
		advancedPrompt := promptui.Prompt{
			Label: "曜日を入力してください (月曜日: 1, 火曜日: 2, ..., 日曜日: 7)",
			Validate: func(input string) error {
				_, err := strconv.Atoi(input)
				if err != nil {
					return err
				}

				return nil
			},
		}

		advancedResult, err := advancedPrompt.Run()
		if err != nil {
			log.Fatalf("曜日入力エラー: %v", err)
		}
		return fmt.Sprintf("on %s", advancedResult)
	}
	if recurringOptions[i].Value == "interval" {
		intervalPrompt := promptui.Prompt{
			Label: "間隔を入力してください (数字)",
			Validate: func(input string) error {
				_, err := strconv.Atoi(input)
				if err != nil {
					return err
				}
				return nil
			},
		}

		intervalResult, err := intervalPrompt.Run()
		if err != nil {
			log.Fatalf("間隔入力エラー: %v", err)
		}
		return fmt.Sprintf("every %s", intervalResult)

	}
	return recurringOptions[i].Value

}
