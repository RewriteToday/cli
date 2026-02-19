package cmd

import (
	"fmt"

	"github.com/RewriteToday/cli/internal/profile"
	"github.com/RewriteToday/cli/internal/render"
	"github.com/RewriteToday/cli/internal/style"
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

	info := style.ProfileInfo{
		Name:   name,
		APIKey: apiKey,
	}

	if format == "json" {
		return style.Print(info, format, noColor)
	}

	printWhoamiText(info, noColor)

	return nil
}

func printWhoamiText(info style.ProfileInfo, noColor bool) {
	fmt.Printf("%s\n", render.Paint("Active profile", render.Bold, noColor))
	fmt.Printf("  %s %s\n", render.Paint("Name:", render.Gray, noColor), render.Paint(info.Name, render.Purple, noColor))
	fmt.Printf("  %s %s\n", render.Paint("API Key:", render.Gray, noColor), render.Paint(maskWhoamiKey(info.APIKey), render.Gray, noColor))
}

func maskWhoamiKey(key string) string {
	if len(key) <= 12 {
		return key + "..."
	}

	return key[:12] + "..."
}

func init() {
	rootCmd.AddCommand(whoamiCmd)
}
