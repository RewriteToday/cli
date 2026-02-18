package cmd

import (
	"github.com/RewriteToday/cli/internal/commands/update"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Aliases: []string{"upgrade"},
	Short: "Update Rewrite CLI to the latest version",
	RunE:  func(cmd *cobra.Command, args []string) error {
		noColor, _ := cmd.Flags().GetBool("no-color")
		
		return update.Update(noColor)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
