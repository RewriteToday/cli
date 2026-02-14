package cmd

import (
	"github.com/rewritestudios/cli/internal/version"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Rewrite CLI to the latest version",
	RunE: func(cmd *cobra.Command, args []string) error {
		noColor, err := cmd.Flags().GetBool("no-color")

		if err != nil {
			return err
		}

		return version.Update(noColor)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
