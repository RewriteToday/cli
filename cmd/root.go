package cmd

import (
	"strings"

	"github.com/RewriteToday/cli/internal/clierr"
	"github.com/RewriteToday/cli/internal/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "rewrite",
	SilenceUsage:  true,
	SilenceErrors: true,
	Version:       version.Version,
	Short:         "A developer-first CLI to integrate Rewrite in your workflow",
	Long: `Rewrite CLI helps you authenticate, manage profiles, trigger test events,
and inspect webhook logs during local development.`,
	Example: `  rewrite login my-profile
  rewrite whoami --output json
  rewrite trigger sms.created -i
  rewrite logs list --limit 50
  rewrite completion zsh > ~/.zsh/completions/_rewrite`,
	PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
		return validateAndNormalizeOutputFormat(cmd)
	},
}

func init() {
	rootCmd.PersistentFlags().Bool("no-color", false, "Remove color from the output")
	rootCmd.PersistentFlags().BoolP("interactive", "i", false, "Run in interactive mode")
	rootCmd.PersistentFlags().StringP("output", "o", "text", "Output format (text or json)")
}

func Execute() error {
	return rootCmd.Execute()
}

func validateAndNormalizeOutputFormat(cmd *cobra.Command) error {
	format, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	normalized, err := normalizeOutputFormat(format)
	if err != nil {
		return err
	}

	return cmd.Flags().Set("output", normalized)
}

var formats = []string{"text", "json"}

func normalizeOutputFormat(raw string) (string, error) {
	format := strings.ToLower(strings.TrimSpace(raw))
	if isSupportedOutputFormat(format) {
		return format, nil
	}

	return "", clierr.Errorf(
		clierr.CodeUsage,
		"invalid output format %q (use one of: %s)",
		raw,
		strings.Join(formats, ", "),
	)
}

func isSupportedOutputFormat(format string) bool {
	for _, supported := range formats {
		if format == supported {
			return true
		}
	}

	return false
}

func ResolveOutputFormat(args []string) string {
	for i := 0; i < len(args); i++ {
		arg := args[i]

		if strings.HasPrefix(arg, "--output=") {
			format, err := normalizeOutputFormat(strings.TrimPrefix(arg, "--output="))
			if err == nil {
				return format
			}
			return "text"
		}

		if arg == "--output" || arg == "-o" {
			if i+1 >= len(args) {
				return "text"
			}

			format, err := normalizeOutputFormat(args[i+1])
			if err == nil {
				return format
			}

			return "text"
		}
	}

	return "text"
}
