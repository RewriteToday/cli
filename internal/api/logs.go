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

func (c *Client) ListLogs(limit int, cursor string) ([]LogEntry, string, error) {
	if limit < 0 {
		limit = 0
	}

	base := time.Now().UTC()
	logs := make([]LogEntry, limit)

	for i := 0; i < limit; i++ {
		eventType := SupportedEvents[i%len(SupportedEvents)]
		payload := MockData(eventType)

		logs[i] = LogEntry{
			ID:        fmt.Sprintf("log_%06d", i+1),
			Timestamp: base.Add(-time.Duration(i*3) * time.Minute).Format(time.RFC3339),
			EventType: string(eventType),
			Status:    eventStatus(eventType),
			Payload:   payload,
		}
	}

	_ = cursor
	return logs, "", nil
}

func eventStatus(eventType EventType) string {
	switch eventType {
	case SMSCreated:
		return "created"
	case SMSSent:
		return "sent"
	case SMSDelivered:
		return "delivered"
	case SMSFailed:
		return "failed"
	default:
		return ""
	}
}
