package profile

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/RewriteToday/cli/internal/config"
	"github.com/zalando/go-keyring"
)

const profileKeyPrefix = "profile:"

func kSet(key, val string) error {
	return keyring.Set(config.KeyringService, key, val)
}

func kGet(key string) (string, error) {
	val, err := keyring.Get(config.KeyringService, key)
	if errors.Is(err, keyring.ErrNotFound) {
		return "", fmt.Errorf("key '%s' not found", key)
	}
	if err != nil {
		return "", err
	}

	return val, nil
}

func kDel(key string) error {
	err := keyring.Delete(config.KeyringService, key)
	if errors.Is(err, keyring.ErrNotFound) {
		return fmt.Errorf("key '%s' not found", key)
	}
	if err != nil {
		return err
	}

	return nil
}

func Save(name, apiKey string) error {
	if err := validateProfileName(name); err != nil {
		return err
	}

	if err := kSet(profileKey(name), apiKey); err != nil {
		return fmt.Errorf("failed to save profile: %w", err)
	}

	profiles, _ := List()

	for _, p := range profiles {
		if p == name {
			return nil
		}
	}

	profiles = append(profiles, name)

	return saveProfileList(profiles)
}

func Get(name string) (string, error) {
	key, err := kGet(profileKey(name))
	if err != nil {
		return "", fmt.Errorf("profile '%s' not found: %w", name, err)
	}

	return key, nil
}

func Delete(name string) error {
	if err := kDel(profileKey(name)); err != nil {
		return fmt.Errorf("failed to delete profile '%s': %w", name, err)
	}

	profiles, _ := List()
	filtered := make([]string, 0, len(profiles))

	for _, p := range profiles {
		if p != name {
			filtered = append(filtered, p)
		}
	}

	if err := saveProfileList(filtered); err != nil {
		return err
	}

	active, _, _ := GetActive()
	if active == name {
		_ = kDel(config.ActiveKey)
	}

	return nil
}

func List() ([]string, error) {
	data, err := kGet(config.ProfilesKey)
	if err != nil {
		return []string{}, nil
	}

	var profiles []string
	if err := json.Unmarshal([]byte(data), &profiles); err != nil {
		return []string{}, nil
	}

	return profiles, nil
}

func Exists(name string) bool {
	_, err := Get(name)
	return err == nil
}

func saveProfileList(profiles []string) error {
	data, err := json.Marshal(profiles)
	if err != nil {
		return fmt.Errorf("failed to serialize profile list: %w", err)
	}

	return kSet(config.ProfilesKey, string(data))
}

func profileKey(name string) string {
	return profileKeyPrefix + name
}

func validateProfileName(name string) error {
	if strings.HasPrefix(name, "__") {
		return fmt.Errorf("profile name cannot start with '__'")
	}
	if strings.HasPrefix(name, profileKeyPrefix) {
		return fmt.Errorf("profile name cannot start with '%s'", profileKeyPrefix)
	}

	return nil
}
