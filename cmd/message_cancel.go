package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	commandmessages "github.com/RewriteToday/cli/internal/commands/messages"
	"github.com/spf13/cobra"
)

var messageCancelCmd = &cobra.Command{
	Use:     "cancel [id]",
	Short:   "Cancel a queued or scheduled message",
	Long:    "Call the `/messages/:id/cancel` endpoint for a message that is still cancelable.",
	Args:    cobra.ExactArgs(1),
	Example: `  rewrite message cancel 1234567890`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return commandmessages.Cancel(commandmessages.CancelOpts{
			RenderOptions: cliutil.ReadRenderOptions(cmd),
			ID:            args[0],
		})
	},
}

func init() {
	messageCmd.AddCommand(messageCancelCmd)
}
