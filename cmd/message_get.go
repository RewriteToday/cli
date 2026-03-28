package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	commandmessages "github.com/RewriteToday/cli/internal/commands/messages"
	"github.com/spf13/cobra"
)

var messageGetCmd = &cobra.Command{
	Use:     "get [id]",
	Short:   "Fetch a single message",
	Long:    "Retrieve one message by ID from the Rewrite `/messages/:id` endpoint.",
	Args:    cobra.ExactArgs(1),
	Example: `  rewrite message get 1234567890`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return commandmessages.Get(commandmessages.GetOpts{
			RenderOptions: cliutil.ReadRenderOptions(cmd),
			ID:            args[0],
		})
	},
}

func init() {
	messageCmd.AddCommand(messageGetCmd)
}
