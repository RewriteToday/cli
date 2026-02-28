package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/commands"
	"github.com/spf13/cobra"
)

var triggerCmd = &cobra.Command{
	Use:     "trigger [event-type]",
	Short:   "Fire realistic test events on demand",
	Long:    "Trigger Rewrite test events instantly to validate integrations, inspect downstream behavior, and debug faster.",
	Aliases: []string{"test"},
	Args:    cobra.MaximumNArgs(1),
	Example: `  rewrite trigger
  rewrite trigger sms.created
  rewrite trigger sms.sent -i
  rewrite trigger sms.delivered --output json`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return commands.Trigger(commands.TriggerOpts{
			Args:                     args,
			InteractiveRenderOptions: cliutil.ReadInteractiveRenderOptions(cmd),
		})
	},
}

func init() {
	rootCmd.AddCommand(triggerCmd)
}
