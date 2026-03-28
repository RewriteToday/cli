package otp

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
	To             string
	Prefix         string
	ExpiresIn      int
	IdempotencyKey string
}

func Create(opts CreateOpts) error {
	if strings.TrimSpace(opts.To) == "" {
		return clierr.Errorf(clierr.CodeUsage, "missing required flag --to")
	}

	body := map[string]any{
		"to": strings.TrimSpace(opts.To),
	}
	if strings.TrimSpace(opts.Prefix) != "" {
		body["prefix"] = strings.TrimSpace(opts.Prefix)
	}
	if opts.ExpiresIn > 0 {
		body["expiresIn"] = opts.ExpiresIn
	}

	client, err := api.New()
	if err != nil {
		return err
	}

	response, err := client.CreateOTP(context.Background(), body, opts.IdempotencyKey)
	if err != nil {
		return err
	}

	return style.Print(response, opts.Format, opts.NoColor)
}
