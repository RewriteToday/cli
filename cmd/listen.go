package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/commands"
	"github.com/spf13/cobra"
)

var listenCmd = &cobra.Command{
	Use:     "listen",
	Short:   "Capture Rewrite events locally in real time",
	Long:    "Start a lightweight local listener to receive webhook events instantly while you build, test, and debug integrations.",
	Aliases: []string{"webhook", "webhooks"},
	Example: `  rewrite listen
  rewrite listen --port 9090
  rewrite listen --output json`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		render := cliutil.ReadRenderOptions(cmd)

		return commands.Listen(commands.ListenOpts{
			Format:  render.Format,
			NoColor: render.NoColor,
			Port:    cliutil.ReadIntFlag(cmd, "port"),
		})
	},
}

func init() {
	listenCmd.Flags().Int("port", 8080, "Choose the local port where Rewrite events should land")

	rootCmd.AddCommand(listenCmd)
}
