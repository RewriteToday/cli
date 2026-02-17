package api

import (
	"fmt"
	"strings"
)

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

func SupportedEventStrings() []string {
	result := make([]string, len(SupportedEvents))
	for i, e := range SupportedEvents {
		result[i] = string(e)
	}
	return result
}

func ValidateEventType(s string) (EventType, error) {
	for _, e := range SupportedEvents {
		if string(e) == s {
			return e, nil
		}
	}

	supported := make([]string, len(SupportedEvents))
	for i, e := range SupportedEvents {
		supported[i] = string(e)
	}

	return "", fmt.Errorf("unsupported event type '%s', supported: %s", s, strings.Join(supported, ", "))
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
