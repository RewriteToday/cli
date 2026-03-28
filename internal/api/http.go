package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/RewriteToday/cli/internal/clierr"
	"github.com/RewriteToday/cli/internal/version"
)

type Cursor struct {
	Next    string `json:"next,omitempty"`
	Prev    string `json:"prev,omitempty"`
	Persist bool   `json:"persist"`
}

type ListResult[T any] struct {
	Items  []T     `json:"items"`
	Cursor *Cursor `json:"cursor,omitempty"`
}

type apiEnvelope struct {
	OK     bool            `json:"ok"`
	Data   json.RawMessage `json:"data"`
	Cursor json.RawMessage `json:"cursor"`
	Error  *apiErrorBody   `json:"error"`
}

type apiErrorBody struct {
	Code     string         `json:"code"`
	Message  string         `json:"message"`
	Detailed map[string]any `json:"detailed"`
}

func (c *Client) request(
	ctx context.Context,
	method string,
	path string,
	query url.Values,
	body any,
	headers map[string]string,
	out any,
) (*Cursor, error) {
	endpoint, err := c.endpoint(path, query)
	if err != nil {
		return nil, err
	}

	var requestBody io.Reader
	if body != nil {
		payload, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("encode request body: %w", err)
		}

		requestBody = bytes.NewReader(payload)
	}

	req, err := http.NewRequestWithContext(ctx, method, endpoint, requestBody)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "rewrite-cli/"+version.Version)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	for key, value := range headers {
		if strings.TrimSpace(value) == "" {
			continue
		}

		req.Header.Set(key, value)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, clierr.Wrap(clierr.CodeNetwork, err)
	}
	defer resp.Body.Close()

	return decodeResponse(resp, out)
}

func (c *Client) endpoint(path string, query url.Values) (string, error) {
	endpoint, err := url.JoinPath(c.BaseURL, strings.TrimPrefix(path, "/"))
	if err != nil {
		return "", fmt.Errorf("join endpoint url: %w", err)
	}

	parsed, err := url.Parse(endpoint)
	if err != nil {
		return "", fmt.Errorf("parse endpoint url: %w", err)
	}

	if len(query) > 0 {
		parsed.RawQuery = query.Encode()
	}

	return parsed.String(), nil
}

func decodeResponse(resp *http.Response, out any) (*Cursor, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, clierr.Wrap(clierr.CodeNetwork, err)
	}

	if len(bytes.TrimSpace(body)) == 0 {
		if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
			return nil, mapAPIError(resp.StatusCode, nil)
		}

		return nil, nil
	}

	var envelope apiEnvelope
	if err := json.Unmarshal(body, &envelope); err != nil {
		return nil, clierr.Wrap(
			clierr.CodeNetwork,
			fmt.Errorf("decode api response: %w", err),
		)
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices || !envelope.OK {
		return nil, mapAPIError(resp.StatusCode, envelope.Error)
	}

	if out != nil && len(envelope.Data) > 0 && string(envelope.Data) != "null" {
		if err := json.Unmarshal(envelope.Data, out); err != nil {
			return nil, fmt.Errorf("decode response payload: %w", err)
		}
	}

	if len(envelope.Cursor) == 0 || string(envelope.Cursor) == "null" {
		return nil, nil
	}

	var cursor Cursor
	if err := json.Unmarshal(envelope.Cursor, &cursor); err != nil {
		return nil, fmt.Errorf("decode response cursor: %w", err)
	}

	return &cursor, nil
}

func mapAPIError(status int, payload *apiErrorBody) error {
	message := "request failed"
	if payload != nil {
		message = strings.TrimSpace(payload.Message)

		if detailedMessage, ok := payload.Detailed["message"].(string); ok && strings.TrimSpace(detailedMessage) != "" {
			message = fmt.Sprintf("%s: %s", message, detailedMessage)
		}

		if message == "" && payload.Code != "" {
			message = payload.Code
		}
	}

	switch status {
	case http.StatusUnauthorized, http.StatusForbidden:
		return clierr.Errorf(clierr.CodeAuthRequired, "%s", message)
	case http.StatusNotFound:
		return clierr.Errorf(clierr.CodeNotFound, "%s", message)
	case http.StatusBadRequest, http.StatusConflict, http.StatusUnprocessableEntity:
		return clierr.Errorf(clierr.CodeUsage, "%s", message)
	default:
		return clierr.Wrap(clierr.CodeInternal, fmt.Errorf("%s", message))
	}
}
