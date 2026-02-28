package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/commands"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:     "login [profile-name]",
	Short:   "Authenticate with Rewrite and get productive fast",
	Long:    "Connect your account, save a reusable profile, and start working with Rewrite in seconds from the command line.",
	Aliases: []string{"auth"},
	Args:    cobra.MaximumNArgs(1),
	Example: `  rewrite login
  rewrite login team-staging
  rewrite login -i`,
	RunE: func(cmd *cobra.Command, args []string) error {
		options := cliutil.ReadInteractiveRenderOptions(cmd)

		return commands.Login(commands.LoginOpts{
			Args:        args,
			Interactive: options.Interactive,
			Format:      options.Format,
			NoColor:     options.NoColor,
		})
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
