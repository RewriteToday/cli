package cmd

import (
	"fmt"

	"github.com/pkg/browser"
	"github.com/rewritestudios/cli/internal/config"
	"github.com/spf13/cobra"
)

var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Open the Rewrite documentation in your browser",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Opening %s\n", config.DocsURL)

		if err := browser.OpenURL(config.DocsURL); err != nil {
			fmt.Printf("Could not open browser. Visit: %s\n", config.DocsURL)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(docsCmd)
}
