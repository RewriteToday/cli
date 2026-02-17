package api

import (
	"fmt"
	"time"
)

type LogEntry struct {
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	EventType string `json:"event_type"`
	Status    string `json:"status"`
	Payload   any    `json:"payload,omitempty"`
}

type logsResponse struct {
	Logs       []LogEntry `json:"logs"`
	NextCursor string     `json:"next_cursor,omitempty"`
}

func (c *Client) ListLogs(limit int, cursor string) ([]LogEntry, string, error) {
	if limit < 0 {
		limit = 0
	}

	eventTypes := []EventType{SMSCreated, SMSSent, SMSDelivered}
	base := time.Now().UTC()
	logs := make([]LogEntry, 0, limit)

	for i := range limit {
		eventType := eventTypes[i%len(eventTypes)]
		payload := MockData(eventType)
		status, _ := payload["status"].(string)

		logs = append(logs, LogEntry{
			ID:        fmt.Sprintf("log_%06d", i+1),
			Timestamp: base.Add(-time.Duration(i*3) * time.Minute).Format(time.RFC3339),
			EventType: string(eventType),
			Status:    status,
			Payload:   payload,
		})
	}

	_ = cursor
	return logs, "", nil
}
