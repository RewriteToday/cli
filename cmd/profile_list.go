package cmd

import (
	"github.com/RewriteToday/cli/internal/commands/profiles"
	"github.com/spf13/cobra"
)

var profileListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all profiles",
	RunE:  func(cmd *cobra.Command, args []string) error {
		format, _ := cmd.Flags().GetString("output")
		noColor, _ := cmd.Flags().GetBool("no-color")
		
		return profiles.List(profiles.ListOpts{
			Format: format,
			NoColor: noColor,
		})
	},
}

func init() {
	profileCmd.AddCommand(profileListCmd)
}
