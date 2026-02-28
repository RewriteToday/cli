package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	commandlogs "github.com/RewriteToday/cli/internal/commands/logs"
	"github.com/spf13/cobra"
)

var logsListCmd = &cobra.Command{
	Use:     "list",
	Short:   "Show the latest Rewrite logs at a glance",
	Long:    "Pull recent log entries so you can spot deliveries, failures, and payloads quickly without opening the dashboard.",
	Aliases: []string{"ls"},
	Example: `  rewrite logs list
  rewrite logs list --limit 50`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		return commandlogs.List(commandlogs.ListOpts{
			RenderOptions: cliutil.ReadRenderOptions(cmd),
			Limit:         cliutil.ReadIntFlag(cmd, "limit"),
		})
	},
}

func init() {
	logsListCmd.Flags().Int("limit", 20, "Set how many recent log entries you want in each snapshot")
	logsCmd.AddCommand(logsListCmd)
}
