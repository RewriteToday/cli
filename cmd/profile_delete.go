package cmd

import (
	"fmt"

	"github.com/rewritestudios/cli/internal/output"
	"github.com/rewritestudios/cli/internal/profile"
	"github.com/rewritestudios/cli/internal/style"
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
	format, _ := cmd.Flags().GetString("output")
	noColor, _ := cmd.Flags().GetBool("no-color")

	name, err := resolveProfileNameToDelete(args, interactive)
	if err != nil {
		return err
	}

	if interactive {
		cancelled, err := confirmDeleteProfile(name, format, noColor)
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

	return output.Print(fmt.Sprintf("Profile '%s' deleted.", name), format, noColor)
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

		name, err = style.SelectString("Select a profile to delete", profiles)
		if err != nil {
			return "", err
		}
	}

	if name == "" {
		return "", fmt.Errorf("profile name required (or use -i for interactive mode)")
	}

	return name, nil
}

func confirmDeleteProfile(name, format string, noColor bool) (bool, error) {
	confirmed, err := style.Confirm(fmt.Sprintf("Delete profile '%s'?", name))
	if err != nil {
		return false, err
	}

	if confirmed {
		return false, nil
	}

	if err := output.Print("Cancelled.", format, noColor); err != nil {
		return false, err
	}

	return true, nil
}

func init() {
	profileCmd.AddCommand(profileDelCmd)
}
