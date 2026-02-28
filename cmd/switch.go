package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/commands/profiles"
	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:     "switch [profile-name]",
	Short:   "Switch profiles instantly and keep momentum",
	Long:    "Change the active Rewrite profile in one command so you can move between projects and environments without breaking flow.",
	Aliases: []string{"use"},
	Args:    cobra.MaximumNArgs(1),
	Example: `  rewrite switch my-profile
  rewrite switch -i`,
	RunE: func(cmd *cobra.Command, args []string) error {
		options := cliutil.ReadInteractiveRenderOptions(cmd)

		return profiles.Switch(profiles.SwitchOpts{
			Args:        args,
			Format:      options.Format,
			NoColor:     options.NoColor,
			Interactive: options.Interactive,
		})
	},
}

func init() {
	rootCmd.AddCommand(switchCmd)
}
