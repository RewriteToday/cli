package otp

import (
	"context"
	"strings"

	"github.com/RewriteToday/cli/internal/api"
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/clierr"
	"github.com/RewriteToday/cli/internal/style"
)

type VerifyOpts struct {
	cliutil.RenderOptions
	ID   string
	To   string
	Code string
}

func Verify(opts VerifyOpts) error {
	if strings.TrimSpace(opts.To) == "" {
		return clierr.Errorf(clierr.CodeUsage, "missing required flag --to")
	}
	if strings.TrimSpace(opts.Code) == "" {
		return clierr.Errorf(clierr.CodeUsage, "missing required flag --code")
	}

	client, err := api.New()
	if err != nil {
		return err
	}

	response, err := client.VerifyOTP(context.Background(), opts.ID, map[string]any{
		"to":   strings.TrimSpace(opts.To),
		"code": strings.TrimSpace(opts.Code),
	})
	if err != nil {
		return err
	}

	return style.Print(response, opts.Format, opts.NoColor)
}
