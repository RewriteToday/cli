package api

import (
	"context"
	"net/http"
	"net/url"
)

type Webhook struct {
	ID        string   `json:"id"`
	Name      *string  `json:"name,omitempty"`
	Secret    string   `json:"secret,omitempty"`
	Events    []string `json:"events"`
	Status    string   `json:"status"`
	Endpoint  string   `json:"endpoint"`
	CreatedAt string   `json:"createdAt"`
}

type WebhookCreateResponse struct {
	ID        string `json:"id"`
	Secret    string `json:"secret,omitempty"`
	CreatedAt string `json:"createdAt"`
}

type WebhookListParams struct {
	Limit  int
	Before string
	After  string
}

func (c *Client) ListWebhooks(
	ctx context.Context,
	params WebhookListParams,
) (*ListResult[Webhook], error) {
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

	var items []Webhook
	cursor, err := c.request(ctx, http.MethodGet, "/webhooks", query, nil, nil, &items)
	if err != nil {
		return nil, err
	}

	return &ListResult[Webhook]{
		Items:  items,
		Cursor: cursor,
	}, nil
}

func (c *Client) GetWebhook(ctx context.Context, id string) (*Webhook, error) {
	var webhook Webhook
	if _, err := c.request(ctx, http.MethodGet, "/webhooks/"+url.PathEscape(id), nil, nil, nil, &webhook); err != nil {
		return nil, err
	}

	return &webhook, nil
}

func (c *Client) CreateWebhook(
	ctx context.Context,
	body map[string]any,
) (*WebhookCreateResponse, error) {
	var response WebhookCreateResponse
	if _, err := c.request(ctx, http.MethodPost, "/webhooks", nil, body, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DeleteWebhook(ctx context.Context, id string) error {
	_, err := c.request(ctx, http.MethodDelete, "/webhooks/"+url.PathEscape(id), nil, nil, nil, nil)
	return err
}
