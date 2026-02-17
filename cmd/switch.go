package cmd

import (
	"fmt"

	"github.com/rewritestudios/cli/internal/profile"
	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch [profile-name]",
	Short: "Switch the active profile",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var name string

		name = args[0]
		if len(args) > 0 {
			return fmt.Errorf("profile name required")
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
