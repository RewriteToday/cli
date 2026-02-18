package cmd

import (
	"github.com/rewritestudios/cli/internal/profile"
	"github.com/rewritestudios/cli/internal/style"
	"github.com/spf13/cobra"
)

var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Show the current active profile",
	RunE:  runWhoamiCommand,
}

func runWhoamiCommand(cmd *cobra.Command, _ []string) error {
	format, _ := cmd.Flags().GetString("output")
	noColor, _ := cmd.Flags().GetBool("no-color")

	name, apiKey, err := profile.GetActive()
	if err != nil {
		return err
	}

	return style.Print(style.ProfileInfo{
		Name:   name,
		APIKey: apiKey,
	}, format, noColor)
}

func init() {
	rootCmd.AddCommand(whoamiCmd)
}
