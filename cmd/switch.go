package cmd

import (
	"fmt"

	"github.com/rewritestudios/cli/internal/output"
	"github.com/rewritestudios/cli/internal/profile"
	"github.com/rewritestudios/cli/internal/style"
	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch [profile-name]",
	Short: "Switch the active profile",
	Args:  cobra.MaximumNArgs(1),
	RunE:  runSwitchCommand,
}

func runSwitchCommand(cmd *cobra.Command, args []string) error {
	interactive, _ := cmd.Flags().GetBool("interactive")
	format, _ := cmd.Flags().GetString("output")
	noColor, _ := cmd.Flags().GetBool("no-color")

	name, err := resolveSwitchProfileName(args, interactive)
	if err != nil {
		return err
	}

	if err := profile.SetActive(name); err != nil {
		return err
	}

	return output.Print(fmt.Sprintf("Switched to profile '%s'", name), format, noColor)
}

func resolveSwitchProfileName(args []string, interactive bool) (string, error) {
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
			return "", fmt.Errorf("no profiles to switch")
		}

		name, err = style.SelectString("Select a profile", profiles)
		if err != nil {
			return "", err
		}
	}

	if name == "" {
		return "", fmt.Errorf("profile name required (or use -i for interactive mode)")
	}

	return name, nil
}

func init() {
	rootCmd.AddCommand(switchCmd)
}
