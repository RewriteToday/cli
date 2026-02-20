package profile

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/RewriteToday/cli/internal/clierr"
	"github.com/RewriteToday/cli/internal/config"
)

func Save(name, apiKey string) error {
	if err := validateProfileName(name); err != nil {
		return err
	}

	if err := KSet(name, apiKey); err != nil {
		return fmt.Errorf("failed to save profile: %w", err)
	}

	profiles, err := List()
	if err != nil {
		return fmt.Errorf("failed to list profiles: %w", err)
	}

	for _, p := range profiles {
		if p == name {
			return nil
		}
	}

	profiles = append(profiles, name)

	return saveProfileList(profiles)
}

func Get(name string) (string, error) {
	key, err := KGet(name)

	if err != nil {
		if errors.Is(err, ErrKeyNotFound) {
			return "", clierr.Errorf(clierr.CodeNotFound, "profile '%s' not found", name)
		}

		return "", fmt.Errorf("failed to load profile '%s': %w", name, err)
	}

	return key, nil
}

func Delete(name string) error {
	if err := KDelete(name); err != nil {
		return fmt.Errorf("failed to delete profile '%s': %w", name, err)
	}

	profiles, err := List()
	if err != nil {
		return fmt.Errorf("failed to list profiles: %w", err)
	}
	filtered := make([]string, 0, len(profiles))

	for _, p := range profiles {
		if p != name {
			filtered = append(filtered, p)
		}
	}

	if err := saveProfileList(filtered); err != nil {
		return err
	}

	active, _, err := GetActive()
	if err != nil && clierr.CodeOf(err) != clierr.CodeAuthRequired {
		return err
	}

	if active == name {
		if err := KDelete(config.ActiveKey); err != nil && !errors.Is(err, ErrKeyNotFound) {
			return fmt.Errorf("failed to clear active profile: %w", err)
		}
	}

	return nil
}

func DeleteAll() (int, error) {
	profiles, err := List()

	if err != nil {
		return 0, err
	}

	for _, name := range profiles {
		if err := KDelete(name); err != nil {
			return 0, fmt.Errorf("failed to delete profile '%s': %w", name, err)
		}
	}

	if err := saveProfileList([]string{}); err != nil {
		return 0, err
	}

	if err := KDelete(config.ActiveKey); err != nil && !errors.Is(err, ErrKeyNotFound) {
		return 0, fmt.Errorf("failed to clear active profile: %w", err)
	}

	return len(profiles), nil
}

func List() ([]string, error) {
	data, err := KGet(config.ProfilesKey)

	if err != nil {
		if errors.Is(err, ErrKeyNotFound) {
			return []string{}, nil
		}

		return nil, fmt.Errorf("failed to read profiles list: %w", err)
	}

	var profiles []string

	if err := json.Unmarshal([]byte(data), &profiles); err != nil {
		return nil, fmt.Errorf("failed to decode profiles list: %w", err)
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

	return KSet(config.ProfilesKey, string(data))
}

func validateProfileName(name string) error {
	if strings.HasPrefix(name, "__") {
		return fmt.Errorf("profile name cannot start with '__'")
	}

	if strings.HasPrefix(name, PROFILE_KEY_PREFIX) {
		return fmt.Errorf("profile name cannot start with '%s'", PROFILE_KEY_PREFIX)
	}

	return nil
}
