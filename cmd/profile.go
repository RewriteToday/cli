package cmd

import "github.com/spf13/cobra"

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Manage profiles",
}

func init() {
	rootCmd.AddCommand(profileCmd)
}
