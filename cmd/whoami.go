package cmd

import (
	"github.com/RewriteToday/cli/internal/commands"
	"github.com/spf13/cobra"
)

var whoamiCmd = &cobra.Command{
	Use:     "whoami",
	Short:   "See which Rewrite profile is active right now",
	Long:    "Confirm the active profile before you trigger events, inspect logs, or switch environments.",
	Aliases: []string{"current"},
	Example: `  rewrite whoami
  rewrite whoami --output json`,
	RunE: func(cmd *cobra.Command, args []string) error {
		format, _ := cmd.Flags().GetString("output")
		noColor, _ := cmd.Flags().GetBool("no-color")

		return commands.Whoami(commands.WhoamiOpts{
			Format:  format,
			NoColor: noColor,
		})
	},
}

func init() {
	rootCmd.AddCommand(whoamiCmd)
}
