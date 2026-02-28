package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/commands/profiles"
	"github.com/spf13/cobra"
)

var profileDelCmd = &cobra.Command{
	Use:     "remove [name]",
	Aliases: []string{"rm", "del", "delete"},
	Short:   "Remove profiles you no longer need",
	Long:    "Clean up saved Rewrite profiles and keep your local setup focused on the accounts and environments that matter.",
	Args:    cobra.MaximumNArgs(1),
	Example: `  rewrite profile remove my-profile
  rewrite profile remove -i`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return profiles.Remove(profiles.RemoveOpts{
			Args:               args,
			InteractiveOptions: cliutil.ReadInteractiveOptions(cmd),
		})
	},
}

func init() {
	profileCmd.AddCommand(profileDelCmd)
}
