package cmd

import (
	"github.com/RewriteToday/cli/internal/commands/profiles"
	"github.com/spf13/cobra"
)

var profileSweepCmd = &cobra.Command{
	Use:     "sweep",
	Short:   "Bulk-clean old profiles in seconds",
	Long:    "Remove older profiles in one pass to keep your Rewrite workspace lean and easy to manage.",
	Aliases: []string{"clean"},
	Example: `  rewrite profile sweep`,
	RunE: func(cmd *cobra.Command, args []string) error {
		noColor, _ := cmd.Flags().GetBool("no-color")

		return profiles.Sweep(profiles.SweepOpts{
			NoColor: noColor,
		})
	},
}

func init() {
	profileCmd.AddCommand(profileSweepCmd)
}
