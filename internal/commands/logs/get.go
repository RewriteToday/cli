package logs

import (
	"context"

	"github.com/RewriteToday/cli/internal/api"
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/style"
)

type GetOpts struct {
	cliutil.RenderOptions
	ID string
}

func Get(opts GetOpts) error {
	client, err := api.New()
	if err != nil {
		return err
	}

	entry, err := client.GetLog(context.Background(), opts.ID)
	if err != nil {
		return err
	}

	return style.Print(entry, opts.Format, opts.NoColor)
}
