package update

import (
	"context"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

const SLUG = "RewriteToday/cli"

func checkUpdate(ctx context.Context, crr string) (release *selfupdate.Release, updated bool, err error) {
	latest, found, err := selfupdate.DetectLatest(SLUG)

	if err != nil {
		return nil, false, err
	}

	if !found || crr == "" {
		return nil, false, nil
	}

	current, err := semver.Parse(crr)

	if err != nil {
		return nil, false, err
	}

	if latest.Version.LTE(current) {
		return nil, false, nil
	}

	return latest, true, nil
}
