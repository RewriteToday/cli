package api

import (
	"context"
	"net/http"
	"net/url"
)

type TemplateVariable struct {
	Name     string `json:"name"`
	Fallback string `json:"fallback,omitempty"`
}

type TemplateTag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Template struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Description *string            `json:"description"`
	Content     string             `json:"content"`
	CreatedAt   string             `json:"createdAt"`
	Variables   []TemplateVariable `json:"variables,omitempty"`
	Tags        []TemplateTag      `json:"tags,omitempty"`
	I18N        map[string]string  `json:"i18n,omitempty"`
}

type TemplateCreateResponse struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}

type TemplateListParams struct {
	Limit    int
	Before   string
	After    string
	WithI18N bool
}

func (c *Client) ListTemplates(
	ctx context.Context,
	params TemplateListParams,
) (*ListResult[Template], error) {
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
	if params.WithI18N {
		query.Set("withi18n", "true")
	}

	var items []Template
	cursor, err := c.request(ctx, http.MethodGet, "/templates", query, nil, nil, &items)
	if err != nil {
		return nil, err
	}

	return &ListResult[Template]{
		Items:  items,
		Cursor: cursor,
	}, nil
}

func (c *Client) GetTemplate(
	ctx context.Context,
	identifier string,
	withI18N bool,
) (*Template, error) {
	query := url.Values{}
	if withI18N {
		query.Set("withi18n", "true")
	}

	var template Template
	if _, err := c.request(ctx, http.MethodGet, "/templates/"+url.PathEscape(identifier), query, nil, nil, &template); err != nil {
		return nil, err
	}

	return &template, nil
}

func (c *Client) CreateTemplate(
	ctx context.Context,
	body map[string]any,
) (*TemplateCreateResponse, error) {
	var response TemplateCreateResponse
	if _, err := c.request(ctx, http.MethodPost, "/templates", nil, body, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DeleteTemplate(ctx context.Context, id string) error {
	_, err := c.request(ctx, http.MethodDelete, "/templates/"+url.PathEscape(id), nil, nil, nil, nil)
	return err
}
