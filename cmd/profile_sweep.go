package cmd

import (
	"github.com/RewriteToday/cli/internal/commands/profiles"
	"github.com/spf13/cobra"
)

var profileSweepCmd = &cobra.Command{
	Use:   "sweep",
	Short: "Sweep all profiles created before",
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
