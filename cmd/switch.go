package cmd

import (
	"fmt"

	"github.com/rewritestudios/cli/internal/profile"
	"github.com/rewritestudios/cli/internal/prompt"
	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch [profile-name]",
	Short: "Switch the active profile",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		interactive, _ := cmd.Flags().GetBool("interactive")
		var name string

		if len(args) > 0 {
			name = args[0]
		}

		if name == "" && interactive {
			profiles, err := profile.List()
			if err != nil {
				return err
			}

			if len(profiles) == 0 {
				return fmt.Errorf("no profiles to switch")
			}

			name, err = prompt.SelectString("Select a profile", profiles)
			if err != nil {
				return err
			}
		}

		if name == "" {
			return fmt.Errorf("profile name required (or use -i for interactive mode)")
		}

		if err := profile.SetActive(name); err != nil {
			return err
		}

		fmt.Printf("Switched to profile '%s'\n", name)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(switchCmd)
}
