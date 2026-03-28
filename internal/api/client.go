package api

import (
	"net/http"
	"strings"

	"github.com/RewriteToday/cli/internal/auth"
	"github.com/RewriteToday/cli/internal/config"
)

type Client struct {
	Profile    string
	BaseURL    string
	httpClient interface {
		Do(*http.Request) (*http.Response, error)
	}
}

func New() (*Client, error) {
	httpClient, name, err := auth.NewAuthenticatedClient()
	if err != nil {
		return nil, err
	}

	return NewWithHTTPClient(name, httpClient, config.APIBaseURL()), nil
}

func NewWithHTTPClient(
	profileName string,
	httpClient interface {
		Do(*http.Request) (*http.Response, error)
	},
	baseURL string,
) *Client {
	return &Client{
		Profile:    profileName,
		BaseURL:    strings.TrimRight(baseURL, "/"),
		httpClient: httpClient,
	}
}
