package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/commands/update"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"upgrade"},
	Short:   "Upgrade to the latest Rewrite CLI in one step",
	Long:    "Pull the newest Rewrite CLI release so you always have the latest fixes, polish, and developer experience improvements.",
	Example: `  rewrite update`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		return update.Update(cliutil.ReadBoolFlag(cmd, "no-color"))
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
