package messages

import (
	"context"
	"fmt"

	"github.com/RewriteToday/cli/internal/api"
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/style"
)

type CancelOpts struct {
	cliutil.RenderOptions
	ID string
}

func Cancel(opts CancelOpts) error {
	client, err := api.New()
	if err != nil {
		return err
	}

	if err := client.CancelMessage(context.Background(), opts.ID); err != nil {
		return err
	}

	return style.Print(
		fmt.Sprintf("Message %s canceled.", opts.ID),
		opts.Format,
		opts.NoColor,
	)
}
