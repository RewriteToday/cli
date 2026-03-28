package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	commandwebhooks "github.com/RewriteToday/cli/internal/commands/webhooks"
	"github.com/spf13/cobra"
)

var webhookDeleteCmd = &cobra.Command{
	Use:     "delete [id]",
	Short:   "Delete a webhook",
	Long:    "Delete a webhook by ID through `/webhooks/:id`.",
	Aliases: []string{"remove", "rm"},
	Args:    cobra.ExactArgs(1),
	Example: `  rewrite webhook delete 1234567890`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return commandwebhooks.Delete(commandwebhooks.DeleteOpts{
			RenderOptions: cliutil.ReadRenderOptions(cmd),
			ID:            args[0],
		})
	},
}

func init() {
	webhookCmd.AddCommand(webhookDeleteCmd)
}
