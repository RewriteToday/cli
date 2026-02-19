package profile

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/RewriteToday/cli/internal/config"
)

func Save(name, apiKey string) error {
	if err := validateProfileName(name); err != nil {
		return err
	}

	if err := KSet(name, apiKey); err != nil {
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
	key, err := KGet(name)
	
	if err != nil {
		return "", fmt.Errorf("profile '%s' not found: %w", name, err)
	}

	return key, nil
}

func Delete(name string) error {
	if err := KDelete(name); err != nil {
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
		_ = KDelete(config.ActiveKey)
	}

	return nil
}

func List() ([]string, error) {
	data, err := KGet(config.ProfilesKey)
	
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
