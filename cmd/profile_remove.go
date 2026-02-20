package cmd

import (
	"github.com/RewriteToday/cli/internal/commands/profiles"
	"github.com/spf13/cobra"
)

var profileDelCmd = &cobra.Command{
	Use:     "remove [name]",
	Aliases: []string{"rm", "del", "delete"},
	Short:   "Remove a logged-in profile",
	Args:    cobra.MaximumNArgs(1),
	Example: `  rewrite profile remove my-profile
  rewrite profile remove -i`,
	RunE: func(cmd *cobra.Command, args []string) error {
		noColor, _ := cmd.Flags().GetBool("no-color")
		interactive, _ := cmd.Flags().GetBool("interactive")

		return profiles.Remove(profiles.RemoveOpts{
			Args:        args,
			NoColor:     noColor,
			Interactive: interactive,
		})
	},
}

func init() {
	profileCmd.AddCommand(profileDelCmd)
}
