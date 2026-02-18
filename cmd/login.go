package cmd

import (
	"fmt"

	"github.com/RewriteToday/cli/internal/auth"
	"github.com/RewriteToday/cli/internal/profile"
	"github.com/RewriteToday/cli/internal/style"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login [profile-name]",
	Short: "Authenticate with Rewrite and save a profile",
	Args:  cobra.MaximumNArgs(1),
	RunE:  runLoginCommand,
}

func runLoginCommand(cmd *cobra.Command, args []string) error {
	interactive, _ := cmd.Flags().GetBool("interactive")
	format, _ := cmd.Flags().GetString("output")
	noColor, _ := cmd.Flags().GetBool("no-color")

	name, err := resolveLoginProfileName(args, interactive)
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
	}, format, noColor)
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

func init() {
	rootCmd.AddCommand(loginCmd)
}
