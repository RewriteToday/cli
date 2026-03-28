package webhooks

import (
	"context"

	"github.com/RewriteToday/cli/internal/api"
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/style"
)

type ListOpts struct {
	cliutil.RenderOptions
	Limit  int
	Before string
	After  string
}

func List(opts ListOpts) error {
	client, err := api.New()
	if err != nil {
		return err
	}

	result, err := client.ListWebhooks(context.Background(), api.WebhookListParams{
		Limit:  opts.Limit,
		Before: opts.Before,
		After:  opts.After,
	})
	if err != nil {
		return err
	}

	return style.Print(result, opts.Format, opts.NoColor)
}
