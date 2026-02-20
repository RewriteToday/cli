package profile

import (
	"errors"
	"fmt"

	"github.com/RewriteToday/cli/internal/clierr"
	"github.com/RewriteToday/cli/internal/config"
)

func GetActive() (string, string, error) {
	name, err := KGet(config.ActiveKey)

	if err != nil {
		if errors.Is(err, ErrKeyNotFound) {
			return "", "", clierr.Errorf(clierr.CodeAuthRequired, "no active profile set, run 'rewrite login' first")
		}

		return "", "", fmt.Errorf("failed to read active profile: %w", err)
	}

	key, err := Get(name)

	if err != nil {
		return "", "", clierr.Wrap(clierr.CodeAuthRequired, fmt.Errorf("active profile '%s' not found: %w", name, err))
	}

	return name, key, nil
}

func SetActive(name string) error {
	if !Exists(name) {
		return clierr.Errorf(clierr.CodeNotFound, "profile '%s' does not exist", name)
	}

	return KSet(config.ActiveKey, name)
}
