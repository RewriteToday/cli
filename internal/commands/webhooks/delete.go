package webhooks

import (
	"context"
	"fmt"

	"github.com/RewriteToday/cli/internal/api"
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/style"
)

type DeleteOpts struct {
	cliutil.RenderOptions
	ID string
}

func Delete(opts DeleteOpts) error {
	client, err := api.New()
	if err != nil {
		return err
	}

	if err := client.DeleteWebhook(context.Background(), opts.ID); err != nil {
		return err
	}

	return style.Print(
		fmt.Sprintf("Webhook %s deleted.", opts.ID),
		opts.Format,
		opts.NoColor,
	)
}
