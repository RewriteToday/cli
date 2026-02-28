package commands

import (
	"fmt"

	"github.com/RewriteToday/cli/internal/auth"
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/profile"
	"github.com/RewriteToday/cli/internal/style"
)

type LoginOpts struct {
	cliutil.InteractiveRenderOptions
	Args []string
}

func Login(opts LoginOpts) error {
	name, err := resolveLoginProfileName(opts.Args, opts.Interactive)
	if err != nil {
		return err
	}

	apiKey, err := auth.RunOAuthFlow()
	if err != nil {
		return fmt.Errorf("authentication failed: %w", err)
	}

	if err := profile.Save(name, apiKey); err != nil {
		return fmt.Errorf("failed to save profile: %w", err)
	}

	if err := profile.SetActive(name); err != nil {
		return fmt.Errorf("failed to set active profile: %w", err)
	}

	return style.Print(style.ProfileInfo{
		Name:   name,
		APIKey: apiKey,
	}, opts.Format, opts.NoColor)
}

func resolveLoginProfileName(args []string, interactive bool) (string, error) {
	if len(args) > 0 {
		return args[0], nil
	}

	if interactive {
		return style.InputString("Profile name", "my-profile")
	}

	return profile.GenerateRandomName(), nil
}
