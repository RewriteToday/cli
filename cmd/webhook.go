package cmd

import "github.com/spf13/cobra"

var webhookCmd = &cobra.Command{
	Use:     "webhook",
	Short:   "Manage Rewrite webhooks",
	Long:    "List, fetch, create, and delete webhooks using the project API-key surface exposed by `/webhooks`.",
	Aliases: []string{"webhooks", "wh"},
	Example: `  rewrite webhook list
  rewrite webhook create --endpoint https://example.com/webhooks --event message.sent`,
}

func init() {
	rootCmd.AddCommand(webhookCmd)
}
