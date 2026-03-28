package cmd

import "github.com/spf13/cobra"

var templateCmd = &cobra.Command{
	Use:     "template",
	Short:   "Manage Rewrite templates",
	Long:    "List, fetch, create, and delete SMS templates through the real `/templates` API.",
	Aliases: []string{"templates", "tpl"},
	Example: `  rewrite template list
  rewrite template create --name welcome --content "Hello {{name}}"`,
}

func init() {
	rootCmd.AddCommand(templateCmd)
}
