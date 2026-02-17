package cmd

import (
	"github.com/rewritestudios/cli/internal/api"
	"github.com/rewritestudios/cli/internal/output"
	"github.com/spf13/cobra"
)

var logsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List recent logs",
	RunE: func(cmd *cobra.Command, args []string) error {
		format, _ := cmd.Flags().GetString("output")
		noColor, _ := cmd.Flags().GetBool("no-color")
		limit, _ := cmd.Flags().GetInt("limit")

		client, err := api.New()
		if err != nil {
			return err
		}

		logs, _, err := client.ListLogs(limit, "")
		if err != nil {
			return err
		}

		items := make([]output.LogEntry, len(logs))
		for i, l := range logs {
			items[i] = output.LogEntry{
				ID:        l.ID,
				Timestamp: l.Timestamp,
				EventType: l.EventType,
				Status:    l.Status,
				Payload:   l.Payload,
			}
		}

		return output.Print(items, format, noColor)
	},
}

func init() {
	logsListCmd.Flags().Int("limit", 20, "Number of log entries to show")
	logsCmd.AddCommand(logsListCmd)
}
