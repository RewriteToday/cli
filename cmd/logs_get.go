package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	commandlogs "github.com/RewriteToday/cli/internal/commands/logs"
	"github.com/spf13/cobra"
)

var logsGetCmd = &cobra.Command{
	Use:     "get [id]",
	Short:   "Fetch one delivery log",
	Long:    "Retrieve a stored delivery log by ID from the Rewrite `/logs/:id` endpoint.",
	Args:    cobra.ExactArgs(1),
	Example: `  rewrite logs get 1234567890`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return commandlogs.Get(commandlogs.GetOpts{
			RenderOptions: cliutil.ReadRenderOptions(cmd),
			ID:            args[0],
		})
	},
}

func init() {
	logsCmd.AddCommand(logsGetCmd)
}
