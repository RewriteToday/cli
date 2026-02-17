package cmd

import (
	"fmt"

	"github.com/rewritestudios/cli/internal/output"
	"github.com/rewritestudios/cli/internal/profile"
	"github.com/rewritestudios/cli/internal/prompt"
	"github.com/spf13/cobra"
)

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Manage profiles",
}

var profileDelCmd = &cobra.Command{
	Use:     "del <name>",
	Aliases: []string{"delete"},
	Short:   "Delete a profile",
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		interactive, _ := cmd.Flags().GetBool("interactive")

		var name string

		if len(args) > 0 {
			name = args[0]
		} else if interactive {
			profiles, err := profile.List()
			if err != nil {
				return err
			}

			if len(profiles) == 0 {
				return fmt.Errorf("no profiles to delete")
			}

			name, err = prompt.SelectString("Select a profile to delete", profiles)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("profile name required (or use -i for interactive mode)")
		}

		if interactive {
			confirmed, err := prompt.Confirm(fmt.Sprintf("Delete profile '%s'?", name))
			if err != nil {
				return err
			}
			if !confirmed {
				fmt.Println("Cancelled.")
				return nil
			}
		}

		if err := profile.Delete(name); err != nil {
			return err
		}

		fmt.Printf("Profile '%s' deleted.\n", name)
		return nil
	},
}

var profileListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all profiles",
	RunE: func(cmd *cobra.Command, args []string) error {
		format, _ := cmd.Flags().GetString("output")
		noColor, _ := cmd.Flags().GetBool("no-color")

		profiles, err := profile.List()
		if err != nil {
			return err
		}

		activeName, _, _ := profile.GetActive()

		items := make([]output.ProfileListItem, len(profiles))
		for i, p := range profiles {
			apiKey, _ := profile.Get(p)
			items[i] = output.ProfileListItem{
				Name:   p,
				APIKey: apiKey,
				Active: p == activeName,
			}
		}

		return output.Print(items, format, noColor)
	},
}

func init() {
	profileCmd.AddCommand(profileDelCmd)
	profileCmd.AddCommand(profileListCmd)
	rootCmd.AddCommand(profileCmd)
}
