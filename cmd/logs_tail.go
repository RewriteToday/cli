package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	commandlogs "github.com/RewriteToday/cli/internal/commands/logs"
	"github.com/spf13/cobra"
)

var logsTailCmd = &cobra.Command{
	Use:     "tail",
	Short:   "Stream incoming Rewrite logs as they happen",
	Long:    "Open a live log stream in your terminal to debug webhook traffic the moment it arrives.",
	Aliases: []string{"stream"},
	Example: `  rewrite logs tail
  rewrite logs tail --port 9090
  rewrite logs tail --output json`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		return commandlogs.Tail(commandlogs.TailOpts{
			RenderOptions: cliutil.ReadRenderOptions(cmd),
			Port:          cliutil.ReadIntFlag(cmd, "port"),
		})
	},
}

func init() {
	logsTailCmd.Flags().Int("port", 8080, "Choose the local port for your live Rewrite log stream")
	logsCmd.AddCommand(logsTailCmd)
}
