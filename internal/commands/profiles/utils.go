package profiles

import (
	"fmt"

	"github.com/RewriteToday/cli/internal/profile"
	"github.com/RewriteToday/cli/internal/style"
)

func resolveName(args []string) (string, error) {
	var name string

	if len(args) > 0 {
		name = args[0]
	}

	if name == "" {
		profiles, err := profile.List()

		if err != nil {
			return "", err
		}

		if len(profiles) == 0 {
			return "", fmt.Errorf("we did not find any profile")
		}

		name, err := style.SelectString("Select a profile", profiles)

		if err != nil {
			return "", err
		}

		if name == "" {
			return "", fmt.Errorf("a profile name is required")
		}
	}

	return name, nil
}
