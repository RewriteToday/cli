package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/commands"
	"github.com/spf13/cobra"
)

var healthCmd = &cobra.Command{
	Use:     "health",
	Short:   "Check whether the Rewrite API is reachable",
	Long:    "Hit the Rewrite health endpoint so you can verify connectivity before sending traffic.",
	Example: `  rewrite health`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		return commands.Health(cliutil.ReadRenderOptions(cmd))
	},
}

func init() {
	rootCmd.AddCommand(healthCmd)
}
