package messages

import (
	"context"

	"github.com/RewriteToday/cli/internal/api"
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/style"
)

type ListOpts struct {
	cliutil.RenderOptions
	Limit   int
	Before  string
	After   string
	Status  string
	Country string
}

func List(opts ListOpts) error {
	client, err := api.New()
	if err != nil {
		return err
	}

	result, err := client.ListMessages(context.Background(), api.MessageListParams{
		Limit:   opts.Limit,
		Before:  opts.Before,
		After:   opts.After,
		Status:  opts.Status,
		Country: opts.Country,
	})
	if err != nil {
		return err
	}

	return style.Print(result, opts.Format, opts.NoColor)
}
