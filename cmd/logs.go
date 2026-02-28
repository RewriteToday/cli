package cmd

import "github.com/spf13/cobra"

var logsCmd = &cobra.Command{
	Use:     "logs",
	Short:   "Inspect live Rewrite logs without leaving your terminal",
	Long:    "List recent deliveries or stream incoming webhook logs in real time to debug faster and stay in flow.",
	Aliases: []string{"log"},
	Example: `  rewrite logs list
  rewrite logs tail`,
}

func init() {
	rootCmd.AddCommand(logsCmd)
}
