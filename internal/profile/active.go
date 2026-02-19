package profile

import (
	"fmt"

	"github.com/RewriteToday/cli/internal/config"
)

func GetActive() (string, string, error) {
	name, err := KGet(config.ActiveKey)

	if err != nil {
		return "", "", fmt.Errorf("no active profile set, run 'rewrite login' first")
	}

	key, err := Get(name)

	if err != nil {
		return "", "", fmt.Errorf("active profile '%s' not found: %w", name, err)
	}

	return name, key, nil
}

func SetActive(name string) error {
	if !Exists(name) {
		return fmt.Errorf("profile '%s' does not exist", name)
	}

	return KSet(config.ActiveKey, name)
}
