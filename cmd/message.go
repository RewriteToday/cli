package cmd

import "github.com/spf13/cobra"

var messageCmd = &cobra.Command{
	Use:     "message",
	Short:   "Send and inspect real Rewrite messages",
	Long:    "Work with the `/messages` API directly from the terminal: send, fetch, list, and cancel messages.",
	Aliases: []string{"messages", "msg"},
	Example: `  rewrite message send --to +5511999999999 --content "Hello"
  rewrite message list --limit 20
  rewrite message get 1234567890`,
}

func init() {
	rootCmd.AddCommand(messageCmd)
}
