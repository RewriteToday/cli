package profiles

import (
	"fmt"

	"github.com/RewriteToday/cli/internal/profile"
	"github.com/RewriteToday/cli/internal/render"
)

type SweepOpts struct {
	NoColor bool
}

func Sweep(opts SweepOpts) error {
	removed, err := profile.DeleteAll()

	if err != nil {
		return err
	}

	if removed == 0 {
		fmt.Println("No profiles found.")

		return nil
	}

	fmt.Printf(
		"Removed %s profile(s) successfully\n",
		render.Paint(fmt.Sprintf("%d", removed), render.Gray, opts.NoColor),
	)

	return nil
}
