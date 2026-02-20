package profiles

import (
	"fmt"

	"github.com/RewriteToday/cli/internal/profile"
	"github.com/RewriteToday/cli/internal/render"
	"github.com/RewriteToday/cli/internal/style"
)

type RemoveOpts struct {
	Args                 []string
	Interactive, NoColor bool
}

func Remove(opts RemoveOpts) error {
	name, err := resolveName(opts.Args, opts.Interactive)

	if err != nil {
		return err
	}

	if opts.Interactive {
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
