package cmd

import "github.com/spf13/cobra"

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "View and stream logs",
}

func init() {
	rootCmd.AddCommand(logsCmd)
}
