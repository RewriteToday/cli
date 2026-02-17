package cmd

import (
	"github.com/rewritestudios/cli/internal/output"
	"github.com/rewritestudios/cli/internal/profile"
	"github.com/spf13/cobra"
)

var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Show the current active profile",
	RunE: func(cmd *cobra.Command, args []string) error {
		format, _ := cmd.Flags().GetString("output")
		noColor, _ := cmd.Flags().GetBool("no-color")

		name, apiKey, err := profile.GetActive()
		if err != nil {
			return err
		}

		return output.Print(output.ProfileInfo{
			Name:   name,
			APIKey: apiKey,
		}, format, noColor)
	},
}

func init() {
	rootCmd.AddCommand(whoamiCmd)
}
