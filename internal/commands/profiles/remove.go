package profiles

import (
	"fmt"

	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/profile"
	"github.com/RewriteToday/cli/internal/render"
	"github.com/RewriteToday/cli/internal/style"
)

type RemoveOpts struct {
	cliutil.InteractiveOptions
	Args []string
}

func Remove(opts RemoveOpts) error {
	interactive := cliutil.ShouldUseInteractive(opts.Args, opts.Interactive)

	name, err := resolveName(opts.Args, interactive)

	if err != nil {
		return err
	}

	if interactive {
		confirmed, err := confirmRemoval(name)

		if err != nil {
			return err
		}

		if !confirmed {
			fmt.Println("Cancelled successfully.")

			return nil
		}
	}

	if err := profile.Delete(name); err != nil {
		return err
	}

	fmt.Printf("Your profile %s was deleted successfully\n", render.Paint(name, render.Gray, opts.NoColor))

	return nil
}

func confirmRemoval(name string) (bool, error) {
	confirmed, err := style.Confirm("Are you sure you want to remove this profile?")

	if err != nil {
		return false, err
	}

	return confirmed, err
}
