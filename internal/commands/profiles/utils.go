package profiles

import (
	"github.com/RewriteToday/cli/internal/clierr"
	"github.com/RewriteToday/cli/internal/profile"
	"github.com/RewriteToday/cli/internal/style"
)

func resolveName(args []string, interactive bool) (string, error) {
	var name string

	if len(args) > 0 {
		name = args[0]
	}

	if name == "" {
		if !interactive {
			return "", clierr.Errorf(clierr.CodeUsage, "profile name required (or use -i for interactive mode)")
		}

		profiles, err := profile.List()

		if err != nil {
			return "", err
		}

		if len(profiles) == 0 {
			return "", clierr.Errorf(clierr.CodeNotFound, "we did not find any profile")
		}

		name, err := style.SelectString("Select a profile", profiles)

		if err != nil {
			return "", err
		}

		if name == "" {
			return "", clierr.Errorf(clierr.CodeUsage, "a profile name is required")
		}
	}

	return name, nil
}
