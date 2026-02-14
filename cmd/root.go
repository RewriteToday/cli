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

func Execute() error {
	rootCmd.PersistentFlags().Bool("no-color", false, "Remove color from the output")

	return rootCmd.Execute()
}
