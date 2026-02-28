package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/commands/profiles"
	"github.com/spf13/cobra"
)

var profileListCmd = &cobra.Command{
	Use:     "list",
	Short:   "See every saved profile in one place",
	Long:    "List your saved Rewrite profiles so switching contexts stays simple and organized.",
	Aliases: []string{"ls"},
	Example: `  rewrite profile list`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		return profiles.List(cliutil.ReadRenderOptions(cmd))
	},
}

func init() {
	profileCmd.AddCommand(profileListCmd)
}
