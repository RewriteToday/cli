package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "rewrite",
	SilenceUsage:  true,
	SilenceErrors: true,
	Version:       version.Version,
	Short:         "The fastest way to test, monitor, and ship Rewrite integrations",
	Long: `Rewrite CLI gives you a faster, cleaner way to build with Rewrite from the terminal.

Authenticate in seconds, switch between profiles, trigger realistic test events,
and inspect live webhook traffic without slowing down your local workflow.`,
	Example: `  rewrite login my-profile
  rewrite whoami --output json
  rewrite trigger sms.created -i
  rewrite logs list --limit 50
  rewrite completion zsh > ~/.zsh/completions/_rewrite`,
	PersistentPreRunE: cliutil.ValidateAndNormalizeOutputFormat,
}

func init() {
	rootCmd.PersistentFlags().Bool("no-color", false, "Render clean, color-free output for piping, logs, and minimal terminals")
	rootCmd.PersistentFlags().BoolP("interactive", "i", false, "Use guided prompts for a faster, more intuitive command flow")
	rootCmd.PersistentFlags().StringP("output", "o", "text", "Choose your output style: text for quick reads or json for automation")
}

func Execute() error {
	configureHelp()

	return rootCmd.Execute()
}
