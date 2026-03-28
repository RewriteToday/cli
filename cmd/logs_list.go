package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	commandlogs "github.com/RewriteToday/cli/internal/commands/logs"
	"github.com/spf13/cobra"
)

var logsListCmd = &cobra.Command{
	Use:     "list",
	Short:   "List delivery logs for a webhook",
	Long:    "Query `/webhooks/:id/logs` so you can inspect delivery attempts, failures, and retries from the terminal.",
	Aliases: []string{"ls"},
	Example: `  rewrite logs list
  rewrite logs list --webhook 1234567890
  rewrite logs list --webhook 1234567890 --status FAILED --limit 50`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		return commandlogs.List(commandlogs.ListOpts{
			RenderOptions: cliutil.ReadRenderOptions(cmd),
			WebhookID:     cliutil.ReadStringFlag(cmd, "webhook"),
			Limit:         cliutil.ReadIntFlag(cmd, "limit"),
			Before:        cliutil.ReadStringFlag(cmd, "before"),
			After:         cliutil.ReadStringFlag(cmd, "after"),
			Type:          cliutil.ReadStringFlag(cmd, "type"),
			Status:        cliutil.ReadStringFlag(cmd, "status"),
		})
	},
}

func init() {
	logsListCmd.Flags().String("webhook", "", "Webhook ID to list logs for")
	logsListCmd.Flags().Int("limit", 20, "Set how many recent log entries you want in each snapshot")
	logsListCmd.Flags().String("before", "", "Cursor for older log entries")
	logsListCmd.Flags().String("after", "", "Cursor for newer log entries")
	logsListCmd.Flags().String("type", "", "Filter by webhook event type")
	logsListCmd.Flags().String("status", "", "Filter by delivery status")
	logsCmd.AddCommand(logsListCmd)
}
