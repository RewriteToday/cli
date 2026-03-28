package api

import (
	"context"
	"net/http"
	"net/url"
)

type OTPCreateResponse struct {
	ID        string `json:"id"`
	To        string `json:"to"`
	Prefix    string `json:"prefix"`
	ExpiresAt string `json:"expiresAt"`
	CreatedAt string `json:"createdAt"`
}

type OTPVerifyResponse struct {
	ID         string `json:"id"`
	Valid      bool   `json:"valid"`
	VerifiedAt string `json:"verifiedAt"`
}

func (c *Client) CreateOTP(
	ctx context.Context,
	body map[string]any,
	idempotencyKey string,
) (*OTPCreateResponse, error) {
	headers := map[string]string{}
	if idempotencyKey != "" {
		headers["Idempotency-Key"] = idempotencyKey
	}

	var response OTPCreateResponse
	if _, err := c.request(ctx, http.MethodPost, "/otp", nil, body, headers, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) VerifyOTP(
	ctx context.Context,
	id string,
	body map[string]any,
) (*OTPVerifyResponse, error) {
	var response OTPVerifyResponse
	if _, err := c.request(ctx, http.MethodPost, "/otp/"+url.PathEscape(id)+"/verify", nil, body, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
