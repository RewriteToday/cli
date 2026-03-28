package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/RewriteToday/cli/internal/auth"
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/clierr"
	"github.com/RewriteToday/cli/internal/profile"
	"github.com/RewriteToday/cli/internal/style"
)

type LoginOpts struct {
	cliutil.InteractiveRenderOptions
	Args   []string
	APIKey string
}

func Login(opts LoginOpts) error {
	name, err := resolveLoginProfileName(opts.Args, opts.Interactive)
	if err != nil {
		return err
	}

	apiKey, err := resolveLoginAPIKey(opts.APIKey, opts.Interactive)
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

func resolveLoginAPIKey(raw string, interactive bool) (string, error) {
	if apiKey, err := auth.ValidateAPIKey(raw); err == nil {
		return apiKey, nil
	}

	if apiKey, err := auth.ValidateAPIKey(os.Getenv("REWRITE_API_KEY")); err == nil {
		return apiKey, nil
	}

	if interactive {
		apiKey, err := style.InputSecret("Rewrite API key")
		if err != nil {
			return "", err
		}

		return auth.ValidateAPIKey(apiKey)
	}

	if strings.TrimSpace(raw) != "" {
		return "", clierr.Errorf(clierr.CodeUsage, "invalid Rewrite API key format")
	}

	return "", clierr.Errorf(
		clierr.CodeUsage,
		"API key required (pass --api-key, set REWRITE_API_KEY, or use -i)",
	)
}
