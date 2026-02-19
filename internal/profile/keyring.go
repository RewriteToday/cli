package profile

import (
	"errors"
	"fmt"

	"github.com/zalando/go-keyring"
)

const KEYRING_SERVICE, PROFILE_KEY_PREFIX = "rewrite-cli", "profile:"

func KSet(key, value string) error {
	return keyring.Set(KEYRING_SERVICE, prefix(key), value)
}

func KGet(key string) (string, error) {
	value, err := keyring.Get(KEYRING_SERVICE, prefix(key))
	if err != nil {
		if errors.Is(err, keyring.ErrNotFound) {
			return "", fmt.Errorf("key '%s' not found", key)
		}

		return "", err
	}

	return value, nil
}

func KDelete(key string) error {
	if err := keyring.Delete(KEYRING_SERVICE, prefix(key)); err != nil {
		if errors.Is(err, keyring.ErrNotFound) {
			return fmt.Errorf("key '%s' not found", key)
		}

		return err
	}

	return nil
}

func prefix(name string) string {
	return PROFILE_KEY_PREFIX + name
}
