package cmd

import (
	"github.com/rewritestudios/cli/internal/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "rewrite",
	SilenceUsage:  true,
	SilenceErrors: true,
	Version:       version.Version,
	Short:         "A developer-first CLI to integrate Rewrite in your workflow",
}

func init() {
	rootCmd.PersistentFlags().Bool("no-color", false, "Remove color from the output")
	rootCmd.PersistentFlags().BoolP("interactive", "i", false, "Run in interactive mode")
	rootCmd.PersistentFlags().StringP("output", "o", "text", "Output format (text or json)")
}

func Execute() error {
	return rootCmd.Execute()
}
