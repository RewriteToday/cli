package cmd

import (
	"github.com/rewritestudios/cli/internal/output"
	"github.com/rewritestudios/cli/internal/profile"
	"github.com/spf13/cobra"
)

var profileListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all profiles",
	RunE:  runProfileListCommand,
}

func runProfileListCommand(cmd *cobra.Command, _ []string) error {
	format, _ := cmd.Flags().GetString("output")
	noColor, _ := cmd.Flags().GetBool("no-color")

	profiles, err := profile.List()
	if err != nil {
		return err
	}

	items := buildProfileListItems(profiles)
	return output.Print(items, format, noColor)
}

func buildProfileListItems(profiles []string) []output.ProfileListItem {
	activeName, _, _ := profile.GetActive()
	items := make([]output.ProfileListItem, len(profiles))

	for i, p := range profiles {
		apiKey, _ := profile.Get(p)
		items[i] = output.ProfileListItem{
			Name:   p,
			APIKey: apiKey,
			Active: p == activeName,
		}
	}

	return items
}

func init() {
	profileCmd.AddCommand(profileListCmd)
}
