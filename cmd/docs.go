package cmd

import (
	"github.com/RewriteToday/cli/internal/config"
	"github.com/RewriteToday/cli/internal/network"
	"github.com/spf13/cobra"
)

var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Open the Rewrite documentation in your browser",
	RunE: func(cmd *cobra.Command, args []string) error {
		noColor, _ := cmd.Flags().GetBool("no-color")

		return network.OpenURL(config.DocsURL, noColor)
	},
}

func init() {
	rootCmd.AddCommand(docsCmd)
}
