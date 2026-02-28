package cli

import (
	"strings"

	"github.com/RewriteToday/cli/internal/clierr"
	"github.com/spf13/cobra"
)

var formats = []string{"text", "json"}

func ValidateAndNormalizeOutputFormat(cmd *cobra.Command, _ []string) error {
	format, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	normalized, err := NormalizeOutputFormat(format)
	if err != nil {
		return err
	}

	return cmd.Flags().Set("output", normalized)
}

func NormalizeOutputFormat(raw string) (string, error) {
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

func ResolveOutputFormat(args []string) string {
	for i := 0; i < len(args); i++ {
		arg := args[i]

		if strings.HasPrefix(arg, "--output=") {
			format, err := NormalizeOutputFormat(strings.TrimPrefix(arg, "--output="))
			if err == nil {
				return format
			}

			return "text"
		}

		if arg == "--output" || arg == "-o" {
			if i+1 >= len(args) {
				return "text"
			}

			format, err := NormalizeOutputFormat(args[i+1])
			if err == nil {
				return format
			}

			return "text"
		}
	}

	return "text"
}

func isSupportedOutputFormat(format string) bool {
	for _, supported := range formats {
		if format == supported {
			return true
		}
	}

	return false
}
