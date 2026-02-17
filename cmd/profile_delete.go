package cmd

import (
	"fmt"

	"github.com/rewritestudios/cli/internal/profile"
	"github.com/rewritestudios/cli/internal/prompt"
	"github.com/spf13/cobra"
)

var profileDelCmd = &cobra.Command{
	Use:     "del <name>",
	Aliases: []string{"del"},
	Short:   "Delete a profile",
	Args:    cobra.MaximumNArgs(1),
	RunE:    runProfileDeleteCommand,
}

func runProfileDeleteCommand(cmd *cobra.Command, args []string) error {
	interactive, _ := cmd.Flags().GetBool("interactive")

	name, err := resolveProfileNameToDelete(args, interactive)
	if err != nil {
		return err
	}

	if interactive {
		cancelled, err := confirmDeleteProfile(name)
		if err != nil {
			return err
		}
		if cancelled {
			return nil
		}
	}

	if err := profile.Delete(name); err != nil {
		return err
	}

	fmt.Printf("Profile '%s' deleted.\n", name)
	return nil
}

func resolveProfileNameToDelete(args []string, interactive bool) (string, error) {
	var name string
	if len(args) > 0 {
		name = args[0]
	}

	if name == "" && interactive {
		profiles, err := profile.List()
		if err != nil {
			return "", err
		}

		if len(profiles) == 0 {
			return "", fmt.Errorf("no profiles to delete")
		}

		name, err = prompt.SelectString("Select a profile to delete", profiles)
		if err != nil {
			return "", err
		}
	}

	if name == "" {
		return "", fmt.Errorf("profile name required (or use -i for interactive mode)")
	}

	return name, nil
}

func confirmDeleteProfile(name string) (bool, error) {
	confirmed, err := prompt.Confirm(fmt.Sprintf("Delete profile '%s'?", name))
	if err != nil {
		return false, err
	}

	if confirmed {
		return false, nil
	}

	fmt.Println("Cancelled.")
	return true, nil
}

func init() {
	profileCmd.AddCommand(profileDelCmd)
}
