package webhooks

import (
	"context"
	"strings"

	"github.com/RewriteToday/cli/internal/api"
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/clierr"
	"github.com/RewriteToday/cli/internal/style"
)

type CreateOpts struct {
	cliutil.RenderOptions
	Name     string
	Endpoint string
	Secret   string
	Events   []string
}

func Create(opts CreateOpts) error {
	if strings.TrimSpace(opts.Endpoint) == "" {
		return clierr.Errorf(clierr.CodeUsage, "missing required flag --endpoint")
	}
	if len(opts.Events) == 0 {
		return clierr.Errorf(clierr.CodeUsage, "at least one --event is required")
	}

	body := map[string]any{
		"endpoint": strings.TrimSpace(opts.Endpoint),
		"events":   opts.Events,
	}
	if strings.TrimSpace(opts.Name) != "" {
		body["name"] = strings.TrimSpace(opts.Name)
	}
	if strings.TrimSpace(opts.Secret) != "" {
		body["secret"] = strings.TrimSpace(opts.Secret)
	}

	client, err := api.New()
	if err != nil {
		return err
	}

	response, err := client.CreateWebhook(context.Background(), body)
	if err != nil {
		return err
	}

	return style.Print(response, opts.Format, opts.NoColor)
}
