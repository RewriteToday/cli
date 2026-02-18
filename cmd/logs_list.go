package cmd

import (
	"github.com/rewritestudios/cli/internal/api"
	"github.com/rewritestudios/cli/internal/style"
	"github.com/spf13/cobra"
)

var logsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List recent logs",
	RunE:  runLogsListCommand,
}

func runLogsListCommand(cmd *cobra.Command, _ []string) error {
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

	items := buildLogEntries(logs)
	return style.Print(items, format, noColor)
}

func buildLogEntries(logs []api.LogEntry) []style.LogEntry {
	items := make([]style.LogEntry, len(logs))
	for i, l := range logs {
		items[i] = style.LogEntry{
			ID:        l.ID,
			Timestamp: l.Timestamp,
			EventType: l.EventType,
			Status:    l.Status,
			Payload:   l.Payload,
		}
	}

	return items
}

func init() {
	logsListCmd.Flags().Int("limit", 20, "Number of log entries to show")
	logsCmd.AddCommand(logsListCmd)
}
