package cmd

import (
	"github.com/RewriteToday/cli/internal/config"
	"github.com/RewriteToday/cli/internal/network"
	"github.com/spf13/cobra"
)

var docsCmd = &cobra.Command{
	Use:     "docs",
	Short:   "Open the Rewrite docs instantly",
	Long:    "Jump straight into the official Rewrite documentation when you need details, examples, or implementation guidance.",
	Aliases: []string{"doc", "documentation"},
	Example: `  rewrite docs`,
	RunE: func(cmd *cobra.Command, args []string) error {
		noColor, _ := cmd.Flags().GetBool("no-color")

		return network.OpenURL(config.DocsURL, noColor)
	},
}

func init() {
	rootCmd.AddCommand(docsCmd)
}
