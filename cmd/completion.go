package cmd

import (
	"os"

	"github.com/RewriteToday/cli/internal/clierr"
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:                   "completion [bash|zsh|fish|powershell]",
	Short:                 "Generate shell completion script",
	Args:                  cobra.ExactArgs(1),
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	DisableFlagsInUseLine: true,
	Example: `  rewrite completion bash > /etc/bash_completion.d/rewrite
  rewrite completion zsh > ~/.zsh/completions/_rewrite
  rewrite completion fish > ~/.config/fish/completions/rewrite.fish
  rewrite completion powershell > rewrite.ps1`,
	RunE: func(cmd *cobra.Command, args []string) error {
		switch args[0] {
		case "bash":
			return rootCmd.GenBashCompletion(os.Stdout)
		case "zsh":
			return rootCmd.GenZshCompletion(os.Stdout)
		case "fish":
			return rootCmd.GenFishCompletion(os.Stdout, true)
		case "powershell":
			return rootCmd.GenPowerShellCompletionWithDesc(os.Stdout)
		default:
			return clierr.Errorf(clierr.CodeUsage, "unsupported shell %q", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
