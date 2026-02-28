package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/commands"
	"github.com/spf13/cobra"
)

var whoamiCmd = &cobra.Command{
	Use:     "whoami",
	Short:   "See which Rewrite profile is active right now",
	Long:    "Confirm the active profile before you trigger events, inspect logs, or switch environments.",
	Aliases: []string{"current"},
	Example: `  rewrite whoami
  rewrite whoami --output json`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		return commands.Whoami(cliutil.ReadRenderOptions(cmd))
	},
}

func init() {
	rootCmd.AddCommand(whoamiCmd)
}
