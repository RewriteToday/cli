package logs

import (
	"context"
	"strings"

	"github.com/RewriteToday/cli/internal/api"
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/clierr"
	"github.com/RewriteToday/cli/internal/style"
)

type ListOpts struct {
	cliutil.RenderOptions
	WebhookID string
	Limit     int
	Before    string
	After     string
	Type      string
	Status    string
}

func List(opts ListOpts) error {
	if strings.TrimSpace(opts.WebhookID) == "" {
		return clierr.Errorf(clierr.CodeUsage, "missing required flag --webhook")
	}

	client, err := api.New()
	if err != nil {
		return err
	}

	logs, err := client.ListWebhookLogs(context.Background(), api.WebhookLogListParams{
		WebhookID: opts.WebhookID,
		Limit:     opts.Limit,
		Before:    opts.Before,
		After:     opts.After,
		Type:      opts.Type,
		Status:    opts.Status,
	})
	if err != nil {
		return err
	}

	return style.Print(logs, opts.Format, opts.NoColor)
}
