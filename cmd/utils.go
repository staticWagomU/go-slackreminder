package cmd

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

func createSelectPrompt(label string, items interface{}) (int, error) {
	prompt := promptui.Select{
		Label: label,
		Items: items,
		IsVimMode: true,
		Templates: &promptui.SelectTemplates{
			Label:    `{{ .Label }}?`,
			Active:   `▸ {{ .Label | cyan }}`,
			Inactive: `  {{ .Label | cyan }}`,
			Selected: `{{ "✔" | green }} {{ .Label | red }}`,
		},
	}

	i, _, err := prompt.Run()
	return i, err
}

func createPrompt(label string, defaultVal string) (string, error) {
	prompt := promptui.Prompt{
		Label:   label,
		Default: defaultVal,
		Templates: &promptui.PromptTemplates{
			Prompt:  "{{ . }} ",
			Valid:   "{{ . | green }} ",
			Invalid: "{{ . | red }} ",
			Success: "{{ . | bold }} ",
		},
	}

	result, err := prompt.Run()
	return result, err
}

func createPrompt2(label string, defaultVal string, validateFunc func(string) error) (string, error) {
	prompt := promptui.Prompt{
		Label:    label,
		Default:  defaultVal,
		Validate: validateFunc,
		Templates: &promptui.PromptTemplates{
			Prompt:  "{{ . }} ",
			Valid:   "{{ . | green }} ",
			Invalid: "{{ . | red }} ",
			Success: "{{ . | bold }} ",
		},
	}

	result, err := prompt.Run()
	return result, err
}

func selectItems(label string, allItems []*item, selectedPos int) ([]*item, error) {
	// Always prepend a "Done" item to the slice if it doesn't
	// already exist.
	const doneID = "Done"
	if len(allItems) > 0 && allItems[0].ID != doneID {
		var items = []*item{
			{
				ID: doneID,
			},
		}

		allItems = append(items, allItems...)
	}

	// Define promptui template
	templates := &promptui.SelectTemplates{
		Label: `{{if .IsSelected}}
					✔
				{{end}} {{ .ID }} - label`,
		Active:   "→ {{if .IsSelected}}✔ {{end}}{{ .ID | cyan }}",
		Inactive: "{{if .IsSelected}}✔ {{end}}{{ .ID | cyan }}",
	}

	prompt := promptui.Select{
		Label:     "Item",
		Items:     allItems,
		Templates: templates,
		Size:      5,
		// Start the cursor at the currently selected index
		CursorPos:    selectedPos,
		HideSelected: true,
	}

	selectionIdx, _, err := prompt.Run()
	if err != nil {
		return nil, fmt.Errorf("prompt failed: %w", err)
	}

	chosenItem := allItems[selectionIdx]

	if chosenItem.ID != doneID {
		// If the user selected something other than "Done",
		// toggle selection on this item and run the function again.
		chosenItem.IsSelected = !chosenItem.IsSelected
		return selectItems(label, allItems, selectionIdx)
	}

	// If the user selected the "Done" item, return
	// all selected items.
	var selectedItems []*item
	for _, i := range allItems {
		if i.IsSelected {
			selectedItems = append(selectedItems, i)
		}
	}
	return selectedItems, nil
}

