package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/commands/profiles"
	"github.com/spf13/cobra"
)

var profileSweepCmd = &cobra.Command{
	Use:     "sweep",
	Short:   "Bulk-clean old profiles in seconds",
	Long:    "Remove older profiles in one pass to keep your Rewrite workspace lean and easy to manage.",
	Aliases: []string{"clean"},
	Example: `  rewrite profile sweep`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		return profiles.Sweep(profiles.SweepOpts{
			NoColor: cliutil.ReadBoolFlag(cmd, "no-color"),
		})
	},
}

func init() {
	profileCmd.AddCommand(profileSweepCmd)
}
