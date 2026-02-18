package auth

import (
	"fmt"
	"net/http"

	"github.com/RewriteToday/cli/internal/profile"
)

type authTransport struct {
	apiKey  string
	wrapped http.RoundTripper
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req = req.Clone(req.Context())
	req.Header.Set("Authorization", "Bearer "+t.apiKey)
	return t.wrapped.RoundTrip(req)
}

func NewAuthenticatedClient() (*http.Client, string, error) {
	name, apiKey, err := profile.GetActive()
	if err != nil {
		return nil, "", fmt.Errorf("not authenticated: %w", err)
	}

	client := &http.Client{
		Transport: &authTransport{
			apiKey:  apiKey,
			wrapped: http.DefaultTransport,
		},
	}

	return client, name, nil
}
