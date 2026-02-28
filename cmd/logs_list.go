package cmd

import (
	"github.com/RewriteToday/cli/internal/api"
	"github.com/RewriteToday/cli/internal/style"
	"github.com/spf13/cobra"
)

var logsListCmd = &cobra.Command{
	Use:     "list",
	Short:   "Show the latest Rewrite logs at a glance",
	Long:    "Pull recent log entries so you can spot deliveries, failures, and payloads quickly without opening the dashboard.",
	Aliases: []string{"ls"},
	Example: `  rewrite logs list
  rewrite logs list --limit 50`,
	RunE: runLogsListCommand,
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
	logsListCmd.Flags().Int("limit", 20, "Set how many recent log entries you want in each snapshot")
	logsCmd.AddCommand(logsListCmd)
}
