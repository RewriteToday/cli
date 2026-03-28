package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	commandwebhooks "github.com/RewriteToday/cli/internal/commands/webhooks"
	"github.com/spf13/cobra"
)

var webhookGetCmd = &cobra.Command{
	Use:     "get [id]",
	Short:   "Fetch one webhook",
	Long:    "Retrieve a single webhook by ID from `/webhooks/:id`.",
	Args:    cobra.ExactArgs(1),
	Example: `  rewrite webhook get 1234567890`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return commandwebhooks.Get(commandwebhooks.GetOpts{
			RenderOptions: cliutil.ReadRenderOptions(cmd),
			ID:            args[0],
		})
	},
}

func init() {
	webhookCmd.AddCommand(webhookGetCmd)
}
