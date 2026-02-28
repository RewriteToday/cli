package profiles

import (
	"fmt"

	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/profile"
	"github.com/RewriteToday/cli/internal/style"
)

type SwitchOpts struct {
	cliutil.InteractiveRenderOptions
	Args []string
}

func Switch(opts SwitchOpts) error {
	interactive := cliutil.ShouldUseInteractive(opts.Args, opts.Interactive)

	name, err := resolveName(opts.Args, interactive)

	if err != nil {
		return err
	}

	if err := profile.SetActive(name); err != nil {
		return err
	}

	return style.Print(fmt.Sprintf("Switched to profile '%s'", name), opts.Format, opts.NoColor)
}
