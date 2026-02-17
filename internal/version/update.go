package version

import (
	"fmt"
	"github.com/rewritestudios/cli/internal/render"
)

func Update(noColor bool) error {
	fmt.Printf(
		"You're already running the latest version of Rewrite (%s)\n",
		render.Paint("v"+Version, render.Gray, noColor),
	)
	return nil
}
