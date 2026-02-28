package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:                   "completion [bash|zsh|fish|powershell]",
	Short:                 "Generate shell completions for a smoother CLI workflow",
	Long:                  "Install completions for your shell and move through Rewrite commands faster with less typing and fewer mistakes.",
	Args:                  cliutil.ValidateCompletionArgs,
	ValidArgs:             cliutil.SupportedCompletionShells,
	DisableFlagsInUseLine: true,
	Example: `  rewrite completion bash > /etc/bash_completion.d/rewrite
  rewrite completion zsh > ~/.zsh/completions/_rewrite
  rewrite completion fish > ~/.config/fish/completions/rewrite.fish
  rewrite completion powershell > rewrite.ps1`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cliutil.RunCompletion(cmd.Root(), args[0])
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
