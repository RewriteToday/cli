package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/commands"
	"github.com/spf13/cobra"
)

var docsCmd = &cobra.Command{
	Use:     "docs",
	Short:   "Open the Rewrite docs instantly",
	Long:    "Jump straight into the official Rewrite documentation when you need details, examples, or implementation guidance.",
	Aliases: []string{"doc", "documentation"},
	Example: `  rewrite docs`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		return commands.Docs(cliutil.ReadBoolFlag(cmd, "no-color"))
	},
}

func init() {
	rootCmd.AddCommand(docsCmd)
}
