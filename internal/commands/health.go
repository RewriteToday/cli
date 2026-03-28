package commands

import (
	"context"

	"github.com/RewriteToday/cli/internal/api"
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/style"
)

func Health(opts cliutil.RenderOptions) error {
	status, err := api.CheckHealth(context.Background())
	if err != nil {
		return err
	}

	return style.Print(status, opts.Format, opts.NoColor)
}
