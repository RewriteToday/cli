package update

import (
	"context"
	"fmt"
	"os"

	"github.com/RewriteToday/cli/internal/render"
	"github.com/RewriteToday/cli/internal/version"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

func Update(noColor bool) error {
	latest, found, err := checkUpdate(context.TODO(), version.Version)

	if err != nil {
		return fmt.Errorf("couldn’t check for updates: %w", err)
	}

	if !found {
		fmt.Printf(
			"You're already on the latest version of Rewrite CLI (%s)!\n",
			render.Paint(version.Version, render.Gray, noColor),
		)

		return nil
	}

	exe, err := os.Executable()

	if err != nil {
		return fmt.Errorf("couldn’t get executable path: %w", err)
	}

	stop := render.Shimmer(context.Background(), fmt.Sprintf("Rewrite found an update (%s). Downloading right now...", latest.Name))

	err = selfupdate.UpdateTo(latest.AssetURL, exe)

	stop()

	if err != nil {
		return err
	}

	fmt.Printf("You're now in the version %s ✨.\n", latest.Version)

	return nil
}
