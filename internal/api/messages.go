package api

import (
	"context"
	"net/http"
	"net/url"
)

type MessageTag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Message struct {
	ID          string       `json:"id"`
	To          string       `json:"to"`
	From        *string      `json:"from,omitempty"`
	Type        string       `json:"type"`
	Status      string       `json:"status"`
	Country     string       `json:"country"`
	Content     string       `json:"content"`
	CreatedAt   string       `json:"createdAt"`
	DeliveredAt *string      `json:"deliveredAt,omitempty"`
	ScheduledAt *string      `json:"scheduledAt,omitempty"`
	TemplateID  *string      `json:"templateId,omitempty"`
	Tags        []MessageTag `json:"tags,omitempty"`
}

type MessageCreateResponse struct {
	ID        string         `json:"id"`
	CreatedAt string         `json:"createdAt"`
	Analysis  map[string]any `json:"analysis"`
}

type MessageListParams struct {
	Limit   int
	Before  string
	After   string
	Status  string
	Country string
}

func (c *Client) ListMessages(
	ctx context.Context,
	params MessageListParams,
) (*ListResult[Message], error) {
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
	if params.Status != "" {
		query.Set("status", params.Status)
	}
	if params.Country != "" {
		query.Set("country", params.Country)
	}

	var items []Message
	cursor, err := c.request(ctx, http.MethodGet, "/messages", query, nil, nil, &items)
	if err != nil {
		return nil, err
	}

	return &ListResult[Message]{
		Items:  items,
		Cursor: cursor,
	}, nil
}

func (c *Client) GetMessage(ctx context.Context, id string) (*Message, error) {
	var message Message
	if _, err := c.request(ctx, http.MethodGet, "/messages/"+url.PathEscape(id), nil, nil, nil, &message); err != nil {
		return nil, err
	}

	return &message, nil
}

func (c *Client) SendMessage(
	ctx context.Context,
	body map[string]any,
	idempotencyKey string,
) (*MessageCreateResponse, error) {
	headers := map[string]string{}
	if idempotencyKey != "" {
		headers["Idempotency-Key"] = idempotencyKey
	}

	var response MessageCreateResponse
	if _, err := c.request(ctx, http.MethodPost, "/messages", nil, body, headers, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CancelMessage(ctx context.Context, id string) error {
	_, err := c.request(ctx, http.MethodPost, "/messages/"+url.PathEscape(id)+"/cancel", nil, nil, nil, nil)
	return err
}
