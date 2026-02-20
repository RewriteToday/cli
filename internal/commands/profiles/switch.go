package profiles

import (
	"fmt"

	"github.com/RewriteToday/cli/internal/profile"
	"github.com/RewriteToday/cli/internal/style"
)

type SwitchOpts struct {
	NoColor, Interactive bool
	Format               string
	Args                 []string
}

func Switch(opts SwitchOpts) error {
	name, err := resolveName(opts.Args, opts.Interactive)

	if err != nil {
		return err
	}

	if err := profile.SetActive(name); err != nil {
		return err
	}

	return style.Print(fmt.Sprintf("Switched to profile '%s'", name), opts.Format, opts.NoColor)
}
