package version

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/creativeprojects/go-selfupdate"
	"github.com/rewritestudios/cli/internal/render"
)

const SLUG = "rewritestudios/cli"

func Update(noColor bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	latest, found, err := selfupdate.DetectLatest(ctx, selfupdate.ParseSlug(SLUG))

	if err != nil {
		return fmt.Errorf("we could not found the latest version (%w)", err)
	}

	if !found || latest.LessOrEqual(Version) {
		fmt.Printf("You'are already running in the latest version of Rewrite (%s)\n", render.Paint(Version, render.Gray, noColor))

		return nil
	}

	exe, err := os.Executable()

	if err != nil {
		return fmt.Errorf("failed to locate executable (%w)", err)
	}

	fmt.Printf("Updating Rewrite from %s to %s...\n", render.Paint(Version, render.Gray, noColor), render.Paint(latest.Version(), render.Gray, noColor))

	if err := selfupdate.UpdateTo(ctx, latest.AssetURL, latest.AssetName, exe); err != nil {
		return fmt.Errorf("update failed (%w)", err)
	}

	fmt.Printf("You'are now on the %s version of Rewrite!\n", render.Paint(latest.Version(), render.Gray, noColor))

	return nil
}
