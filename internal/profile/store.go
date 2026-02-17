package profile

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/rewritestudios/cli/internal/config"
)

var (
	mu    = sync.Mutex{}
	store = map[string]string{
		"cosmic-falcon":    "rw_live_demo_key_abc123xyz789",
		config.ProfilesKey: `["cosmic-falcon"]`,
		config.ActiveKey:   "cosmic-falcon",
	}
)

func kSet(key, val string) error {
	mu.Lock()
	defer mu.Unlock()
	store[key] = val
	return nil
}

func kGet(key string) (string, error) {
	mu.Lock()
	defer mu.Unlock()

	val, ok := store[key]
	if !ok {
		return "", fmt.Errorf("key '%s' not found", key)
	}

	return val, nil
}

func kDel(key string) error {
	mu.Lock()
	defer mu.Unlock()

	if _, ok := store[key]; !ok {
		return fmt.Errorf("key '%s' not found", key)
	}

	delete(store, key)
	return nil
}

func Save(name, apiKey string) error {
	if strings.HasPrefix(name, "__") {
		return fmt.Errorf("profile name cannot start with '__'")
	}

	if err := kSet(name, apiKey); err != nil {
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
	key, err := kGet(name)
	if err != nil {
		return "", fmt.Errorf("profile '%s' not found: %w", name, err)
	}

	return key, nil
}

func Delete(name string) error {
	if err := kDel(name); err != nil {
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
