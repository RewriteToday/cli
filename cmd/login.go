package cmd

import (
	"fmt"

	"github.com/rewritestudios/cli/internal/auth"
	"github.com/rewritestudios/cli/internal/output"
	"github.com/rewritestudios/cli/internal/profile"
	"github.com/rewritestudios/cli/internal/prompt"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login [profile-name]",
	Short: "Authenticate with Rewrite and save a profile",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		interactive, _ := cmd.Flags().GetBool("interactive")
		format, _ := cmd.Flags().GetString("output")

		var name string

		name = args[0]
		if interactive {
			var err error

			name, err = prompt.InputString("Profile name", "my-profile")
			if err != nil {
				return err
			}
		}

		if len(args) == 0 {
			name = profile.GenerateRandomName()
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

		fmt.Printf("Logged in as '%s'\n", name)

		noColor, _ := cmd.Flags().GetBool("no-color")
		return output.Print(output.ProfileInfo{
			Name:   name,
			APIKey: apiKey,
		}, format, noColor)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
