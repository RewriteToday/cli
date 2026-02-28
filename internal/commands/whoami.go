package commands

import (
	"fmt"

	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/profile"
	"github.com/RewriteToday/cli/internal/render"
	"github.com/RewriteToday/cli/internal/style"
)

func Whoami(opts cliutil.RenderOptions) error {
	name, apiKey, err := profile.GetActive()

	if err != nil {
		return err
	}

	info := style.ProfileInfo{
		Name:   name,
		APIKey: apiKey,
	}

	if opts.Format == "json" {
		return style.Print(info, opts.Format, opts.NoColor)
	}

	printWhoamiText(info, opts.NoColor)

	return nil
}

func printWhoamiText(info style.ProfileInfo, noColor bool) {
	fmt.Printf("%s\n", render.Paint("Active profile", render.Bold, noColor))
	fmt.Printf("  %s %s\n", render.Paint("Name:", render.Gray, noColor), render.Paint(info.Name, render.Purple, noColor))
	fmt.Printf("  %s %s\n", render.Paint("API Key:", render.Gray, noColor), render.Paint(style.MaskKey(info.APIKey), render.Gray, noColor))
}
