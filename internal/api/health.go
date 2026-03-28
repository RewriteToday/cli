package api

import (
	"context"
	"net/http"
	"time"

	"github.com/RewriteToday/cli/internal/config"
)

type HealthStatus struct {
	Uptime int `json:"uptime"`
}

func CheckHealth(ctx context.Context) (*HealthStatus, error) {
	client := NewWithHTTPClient("public", &http.Client{
		Timeout: 15 * time.Second,
	}, config.APIBaseURL())

	var health HealthStatus
	if _, err := client.request(ctx, http.MethodGet, "/health", nil, nil, nil, &health); err != nil {
		return nil, err
	}

	return &health, nil
}
