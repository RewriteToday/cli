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

	activeName, _, err := profile.GetActive()
	if err != nil {
		activeName = ""
	}

	items, err := createProfileListItems(profiles)
	if err != nil {
		return err
	}

	if opts.Format == "json" {
		jsonItems := createProfileListJSONItems(items, activeName)

		return style.Print(jsonItems, opts.Format, opts.NoColor)
	}

	return style.Print(style.ProfileListText{
		ActiveName: activeName,
		Items:      items,
	}, opts.Format, opts.NoColor)
}

func createProfileListItems(profiles []string) ([]style.ProfileListItem, error) {
	items := make([]style.ProfileListItem, len(profiles))

	for i, p := range profiles {
		apiKey, err := profile.Get(p)
		if err != nil {
			return nil, err
		}

		items[i] = style.ProfileListItem{
			Name:   p,
			APIKey: apiKey,
		}
	}

	return items, nil
}

func createProfileListJSONItems(items []style.ProfileListItem, activeName string) []style.ProfileListItemJSON {
	jsonItems := make([]style.ProfileListItemJSON, len(items))

	for i, item := range items {
		jsonItems[i] = style.ProfileListItemJSON{
			Name:         item.Name,
			APIKeyMasked: style.MaskKey(item.APIKey),
			Active:       item.Name == activeName,
		}
	}

	return jsonItems
}
