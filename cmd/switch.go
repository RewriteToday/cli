package cmd

import (
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
		format, _ := cmd.Flags().GetString("output")
		noColor, _ := cmd.Flags().GetBool("no-color")
		interactive, _ := cmd.Flags().GetBool("interactive")

		return profiles.Switch(profiles.SwitchOpts{
			Args:        args,
			Format:      format,
			NoColor:     noColor,
			Interactive: interactive,
		})
	},
}

func init() {
	rootCmd.AddCommand(switchCmd)
}
