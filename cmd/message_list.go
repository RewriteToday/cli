package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	commandmessages "github.com/RewriteToday/cli/internal/commands/messages"
	"github.com/spf13/cobra"
)

var messageListCmd = &cobra.Command{
	Use:   "list",
	Short: "List project messages",
	Long:  "Query the Rewrite `/messages` endpoint with optional cursor and status filters.",
	Example: `  rewrite message list
  rewrite message list --status SENT --limit 50
  rewrite message list --before 1234567890`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		return commandmessages.List(commandmessages.ListOpts{
			RenderOptions: cliutil.ReadRenderOptions(cmd),
			Limit:         cliutil.ReadIntFlag(cmd, "limit"),
			Before:        cliutil.ReadStringFlag(cmd, "before"),
			After:         cliutil.ReadStringFlag(cmd, "after"),
			Status:        cliutil.ReadStringFlag(cmd, "status"),
			Country:       cliutil.ReadStringFlag(cmd, "country"),
		})
	},
}

func init() {
	messageListCmd.Flags().Int("limit", 20, "Maximum number of messages to return")
	messageListCmd.Flags().String("before", "", "Cursor for older messages")
	messageListCmd.Flags().String("after", "", "Cursor for newer messages")
	messageListCmd.Flags().String("status", "", "Filter by message status")
	messageListCmd.Flags().String("country", "", "Filter by ISO 3166-1 alpha-2 country code")

	messageCmd.AddCommand(messageListCmd)
}
