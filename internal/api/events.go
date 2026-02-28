package api

import "github.com/RewriteToday/cli/internal/clierr"

type EventType string

const (
	SMSCreated   EventType = "sms.created"
	SMSSent      EventType = "sms.sent"
	SMSDelivered EventType = "sms.delivered"
)

var SupportedEvents = []EventType{
	SMSCreated,
	SMSSent,
	SMSDelivered,
}

const supportedEventList = "sms.created, sms.sent, sms.delivered"

func SupportedEventStrings() []string {
	result := make([]string, len(SupportedEvents))
	for i, e := range SupportedEvents {
		result[i] = string(e)
	}
	return result
}

func ValidateEventType(s string) (EventType, error) {
	switch EventType(s) {
	case SMSCreated, SMSSent, SMSDelivered:
		return EventType(s), nil
	}

	return "", clierr.Errorf(clierr.CodeUsage, "unsupported event type '%s', supported: %s", s, supportedEventList)
}

func MockData(eventType EventType) map[string]any {
	switch eventType {
	case SMSCreated:
		return map[string]any{
			"id":         "msg_01abc123",
			"to":         "+5511999999999",
			"from":       "+5511888888888",
			"body":       "Hello from Rewrite!",
			"status":     "created",
			"created_at": "2025-01-15T10:30:00Z",
		}
	case SMSSent:
		return map[string]any{
			"id":         "msg_01abc123",
			"to":         "+5511999999999",
			"from":       "+5511888888888",
			"body":       "Hello from Rewrite!",
			"status":     "sent",
			"sent_at":    "2025-01-15T10:30:05Z",
			"created_at": "2025-01-15T10:30:00Z",
		}
	case SMSDelivered:
		return map[string]any{
			"id":           "msg_01abc123",
			"to":           "+5511999999999",
			"from":         "+5511888888888",
			"body":         "Hello from Rewrite!",
			"status":       "delivered",
			"delivered_at": "2025-01-15T10:30:10Z",
			"sent_at":      "2025-01-15T10:30:05Z",
			"created_at":   "2025-01-15T10:30:00Z",
		}
	default:
		return map[string]any{}
	}
}
