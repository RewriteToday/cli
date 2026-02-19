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

	activeName, _, _ := profile.GetActive()
	items := createProfileListItems(profiles)

	if opts.Format == "json" {
		jsonItems := createProfileListJSONItems(items, activeName)

		return style.Print(jsonItems, opts.Format, opts.NoColor)
	}

	return style.Print(style.ProfileListText{
		ActiveName: activeName,
		Items:      items,
	}, opts.Format, opts.NoColor)
}

func createProfileListItems(profiles []string) []style.ProfileListItem {
	items := make([]style.ProfileListItem, len(profiles))

	for i, p := range profiles {
		apiKey, _ := profile.Get(p)

		items[i] = style.ProfileListItem{
			Name:   p,
			APIKey: apiKey,
		}
	}

	return items
}

func createProfileListJSONItems(items []style.ProfileListItem, activeName string) []style.ProfileListItemJSON {
	jsonItems := make([]style.ProfileListItemJSON, len(items))

	for i, item := range items {
		jsonItems[i] = style.ProfileListItemJSON{
			Name:   item.Name,
			APIKey: item.APIKey,
			Active: item.Name == activeName,
		}
	}

	return jsonItems
}
