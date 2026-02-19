package cmd

import (
	"github.com/RewriteToday/cli/internal/commands/profiles"
	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch [profile-name]",
	Short: "Switch the active profile",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		format, _ := cmd.Flags().GetString("output")
		noColor, _ := cmd.Flags().GetBool("no-color")

		return profiles.Switch(profiles.SwitchOpts{
			Args:    args,
			Format:  format,
			NoColor: noColor,
		})
	},
}

func init() {
	rootCmd.AddCommand(switchCmd)
}
