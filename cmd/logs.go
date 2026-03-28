package cmd

import "github.com/spf13/cobra"

var logsCmd = &cobra.Command{
	Use:     "logs",
	Short:   "Inspect Rewrite delivery logs",
	Long:    "Fetch stored delivery logs from the API or stream local webhook traffic in real time while debugging integrations.",
	Aliases: []string{"log"},
	Example: `  rewrite logs list
  rewrite logs get 1234567890
  rewrite logs tail`,
}

func init() {
	rootCmd.AddCommand(logsCmd)
}
