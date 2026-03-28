package templates

import (
	"context"

	"github.com/RewriteToday/cli/internal/api"
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/style"
)

type ListOpts struct {
	cliutil.RenderOptions
	Limit    int
	Before   string
	After    string
	WithI18N bool
}

func List(opts ListOpts) error {
	client, err := api.New()
	if err != nil {
		return err
	}

	result, err := client.ListTemplates(context.Background(), api.TemplateListParams{
		Limit:    opts.Limit,
		Before:   opts.Before,
		After:    opts.After,
		WithI18N: opts.WithI18N,
	})
	if err != nil {
		return err
	}

	return style.Print(result, opts.Format, opts.NoColor)
}
