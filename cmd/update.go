package cmd

import (
	"github.com/RewriteToday/cli/internal/commands/update"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Rewrite CLI to the latest version",
	RunE:  runUpdate,
}

func runUpdate(cmd *cobra.Command, _ []string) error {
	noColor, err := cmd.Flags().GetBool("no-color")

	if err != nil {
		return err
	}

	return update.Update(noColor)
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
