package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	commandwebhooks "github.com/RewriteToday/cli/internal/commands/webhooks"
	"github.com/spf13/cobra"
)

var webhookListCmd = &cobra.Command{
	Use:   "list",
	Short: "List project webhooks",
	Long:  "Query the Rewrite `/webhooks` endpoint with cursor options.",
	Example: `  rewrite webhook list
  rewrite webhook list --limit 50`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		return commandwebhooks.List(commandwebhooks.ListOpts{
			RenderOptions: cliutil.ReadRenderOptions(cmd),
			Limit:         cliutil.ReadIntFlag(cmd, "limit"),
			Before:        cliutil.ReadStringFlag(cmd, "before"),
			After:         cliutil.ReadStringFlag(cmd, "after"),
		})
	},
}

func init() {
	webhookListCmd.Flags().Int("limit", 20, "Maximum number of webhooks to return")
	webhookListCmd.Flags().String("before", "", "Cursor for older webhooks")
	webhookListCmd.Flags().String("after", "", "Cursor for newer webhooks")

	webhookCmd.AddCommand(webhookListCmd)
}
