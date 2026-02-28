package cmd

import "github.com/spf13/cobra"

var profileCmd = &cobra.Command{
	Use:     "profile",
	Short:   "Manage Rewrite profiles with less friction",
	Long:    "Create, inspect, switch, and clean up profiles so moving between environments stays fast and reliable.",
	Aliases: []string{"pf", "profiles"},
	Example: `  rewrite profile list
  rewrite profile remove my-profile`,
}

func init() {
	rootCmd.AddCommand(profileCmd)
}
