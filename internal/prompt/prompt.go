package prompt

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

func InputString(title, placeholder string) (string, error) {
	var value string

	err := huh.NewInput().
		Title(title).
		Placeholder(placeholder).
		Value(&value).
		Run()

	if err != nil {
		return "", fmt.Errorf("input cancelled: %w", err)
	}

	return value, nil
}

func SelectString(title string, options []string) (string, error) {
	var value string

	opts := make([]huh.Option[string], len(options))
	for i, o := range options {
		opts[i] = huh.NewOption(o, o)
	}

	err := huh.NewSelect[string]().
		Title(title).
		Options(opts...).
		Value(&value).
		Run()

	if err != nil {
		return "", fmt.Errorf("selection cancelled: %w", err)
	}

	return value, nil
}

func Confirm(title string) (bool, error) {
	var value bool

	err := huh.NewConfirm().
		Title(title).
		Value(&value).
		Run()

	if err != nil {
		return false, fmt.Errorf("confirmation cancelled: %w", err)
	}

	return value, nil
}

func TriggerEventForm(eventType string) (map[string]any, error) {
	var to, from, body string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("To (phone number)").
				Placeholder("+5511999999999").
				Value(&to),
			huh.NewInput().
				Title("From (phone number)").
				Placeholder("+5511888888888").
				Value(&from),
			huh.NewInput().
				Title("Body (message content)").
				Placeholder("Hello from Rewrite!").
				Value(&body),
		),
	)

	if err := form.Run(); err != nil {
		return nil, fmt.Errorf("form cancelled: %w", err)
	}

	data := map[string]any{
		"to":   to,
		"from": from,
		"body": body,
	}

	return data, nil
}
