package logs

import (
	"github.com/RewriteToday/cli/internal/api"
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/style"
)

type ListOpts struct {
	cliutil.RenderOptions
	Limit int
}

func List(opts ListOpts) error {
	client, err := api.New()
	if err != nil {
		return err
	}

	logs, _, err := client.ListLogs(opts.Limit, "")
	if err != nil {
		return err
	}

	return style.Print(buildLogEntries(logs), opts.Format, opts.NoColor)
}

func buildLogEntries(entries []api.LogEntry) []style.LogEntry {
	items := make([]style.LogEntry, len(entries))

	for i, entry := range entries {
		items[i] = style.LogEntry{
			ID:        entry.ID,
			Timestamp: entry.Timestamp,
			EventType: entry.EventType,
			Status:    entry.Status,
			Payload:   entry.Payload,
		}
	}

	return items
}
