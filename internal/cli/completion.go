package cli

import (
	"os"

	"github.com/RewriteToday/cli/internal/clierr"
	"github.com/spf13/cobra"
)

type completionGenerator func(*cobra.Command) error

var SupportedCompletionShells = []string{"bash", "zsh", "fish", "powershell"}

var completionGenerators = map[string]completionGenerator{
	"bash": func(root *cobra.Command) error {
		return root.GenBashCompletion(os.Stdout)
	},
	"zsh": func(root *cobra.Command) error {
		return root.GenZshCompletion(os.Stdout)
	},
	"fish": func(root *cobra.Command) error {
		return root.GenFishCompletion(os.Stdout, true)
	},
	"powershell": func(root *cobra.Command) error {
		return root.GenPowerShellCompletionWithDesc(os.Stdout)
	},
}

func ValidateCompletionArgs(_ *cobra.Command, args []string) error {
	if err := cobra.ExactArgs(1)(nil, args); err != nil {
		return err
	}

	if _, ok := completionGenerators[args[0]]; ok {
		return nil
	}

	return clierr.Errorf(clierr.CodeUsage, "unsupported shell %q", args[0])
}

func RunCompletion(root *cobra.Command, shell string) error {
	return completionGenerators[shell](root)
}
