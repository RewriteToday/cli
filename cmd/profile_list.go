package cmd

import (
	"github.com/RewriteToday/cli/internal/commands/profiles"
	"github.com/spf13/cobra"
)

var profileListCmd = &cobra.Command{
	Use:     "list",
	Short:   "See every saved profile in one place",
	Long:    "List your saved Rewrite profiles so switching contexts stays simple and organized.",
	Aliases: []string{"ls"},
	Example: `  rewrite profile list`,
	RunE: func(cmd *cobra.Command, args []string) error {
		format, _ := cmd.Flags().GetString("output")
		noColor, _ := cmd.Flags().GetBool("no-color")

		return profiles.List(profiles.ListOpts{
			Format:  format,
			NoColor: noColor,
		})
	},
}

func init() {
	profileCmd.AddCommand(profileListCmd)
}
