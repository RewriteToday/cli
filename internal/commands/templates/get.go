package templates

import (
	"context"

	"github.com/RewriteToday/cli/internal/api"
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/style"
)

type GetOpts struct {
	cliutil.RenderOptions
	Identifier string
	WithI18N   bool
}

func Get(opts GetOpts) error {
	client, err := api.New()
	if err != nil {
		return err
	}

	template, err := client.GetTemplate(context.Background(), opts.Identifier, opts.WithI18N)
	if err != nil {
		return err
	}

	return style.Print(template, opts.Format, opts.NoColor)
}
