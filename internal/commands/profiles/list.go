package profiles

import (
	"github.com/RewriteToday/cli/internal/profile"
	"github.com/RewriteToday/cli/internal/style"
)

type ListOpts struct {
	NoColor bool
	Format  string
}

func List(opts ListOpts) error {
	profiles, err := profile.List()

	if err != nil {
		return err
	}

	items := createProfileListItems(profiles)

	return style.Print(items, opts.Format, opts.NoColor)
}

func createProfileListItems(profiles []string) []style.ProfileListItem {
	activeName, _, _ := profile.GetActive()

	items := make([]style.ProfileListItem, len(profiles))

	for i, p := range profiles {
		apiKey, _ := profile.Get(p)

		items[i] = style.ProfileListItem{
			Name:   p,
			APIKey: apiKey,
			Active: p == activeName,
		}
	}

	return items
}
