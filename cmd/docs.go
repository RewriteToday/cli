package cmd

import (
	"fmt"

	"github.com/pkg/browser"
	"github.com/RewriteToday/cli/internal/config"
	"github.com/RewriteToday/cli/internal/style"
	"github.com/spf13/cobra"
)

var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Open the Rewrite documentation in your browser",
	RunE:  runDocsCommand,
}

func runDocsCommand(cmd *cobra.Command, _ []string) error {
	format, _ := cmd.Flags().GetString("output")
	noColor, _ := cmd.Flags().GetBool("no-color")

	if err := style.Print(fmt.Sprintf("Opening %s", config.DocsURL), format, noColor); err != nil {
		return err
	}

	if err := browser.OpenURL(config.DocsURL); err != nil {
		return style.Print(fmt.Sprintf("Could not open browser. Visit: %s", config.DocsURL), format, noColor)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(docsCmd)
}
