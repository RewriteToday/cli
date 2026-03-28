package api

import (
	"context"
	"net/http"
	"net/url"
)

type LogEntry struct {
	ID        string         `json:"id"`
	URL       string         `json:"url,omitempty"`
	Timestamp string         `json:"timestamp,omitempty"`
	EventType string         `json:"event_type,omitempty"`
	Type      string         `json:"type,omitempty"`
	Status    string         `json:"status"`
	Payload   map[string]any `json:"payload,omitempty"`
	Error     string         `json:"error,omitempty"`
	RetryAt   string         `json:"retryAt,omitempty"`
	CreatedAt string         `json:"createdAt,omitempty"`
	WebhookID *string        `json:"webhookId,omitempty"`
	MessageID *string        `json:"messageId,omitempty"`
	Code      *int           `json:"code,omitempty"`
	Latency   *int           `json:"latency,omitempty"`
	Attempt   *int           `json:"attempt,omitempty"`
}

type WebhookLogListParams struct {
	WebhookID string
	Limit     int
	Before    string
	After     string
	Type      string
	Status    string
}

func (c *Client) GetLog(ctx context.Context, id string) (*LogEntry, error) {
	var entry LogEntry
	if _, err := c.request(ctx, http.MethodGet, "/logs/"+url.PathEscape(id), nil, nil, nil, &entry); err != nil {
		return nil, err
	}

	return &entry, nil
}

func (c *Client) ListWebhookLogs(
	ctx context.Context,
	params WebhookLogListParams,
) (*ListResult[LogEntry], error) {
	query := url.Values{}
	if params.Limit > 0 {
		query.Set("limit", intString(params.Limit))
	}
	if params.Before != "" {
		query.Set("before", params.Before)
	}
	if params.After != "" {
		query.Set("after", params.After)
	}
	if params.Type != "" {
		query.Set("type", params.Type)
	}
	if params.Status != "" {
		query.Set("status", params.Status)
	}

	var items []LogEntry
	cursor, err := c.request(
		ctx,
		http.MethodGet,
		"/webhooks/"+url.PathEscape(params.WebhookID)+"/logs",
		query,
		nil,
		nil,
		&items,
	)
	if err != nil {
		return nil, err
	}

	return &ListResult[LogEntry]{
		Items:  items,
		Cursor: cursor,
	}, nil
}
