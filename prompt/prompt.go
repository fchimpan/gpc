package prompt

import (
	"errors"

	"github.com/manifoldco/promptui"
)

func SelectConfig(configs []string) (string, error) {
	prompt := promptui.Select{
		Label: "Select Config",
		Items: configs,
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}

func InputPageTitle() (string, error) {
	prompt := promptui.Prompt{
		Label:    "Page title",
		Validate: validateEmptyInput("Page title"),
	}
	result, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return result, nil
}

func InputPageBody() (string, error) {
	prompt := promptui.Prompt{
		Label:    "Page body",
		Validate: validateEmptyInput("Page body"),
	}
	result, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return result, nil
}

func validateEmptyInput(label string) func(input string) error {
	return func(input string) error {
		if input == "" {
			return errors.New("must provide a " + label)
		}
		return nil
	}
}
