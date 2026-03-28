package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	commandtemplates "github.com/RewriteToday/cli/internal/commands/templates"
	"github.com/spf13/cobra"
)

var templateListCmd = &cobra.Command{
	Use:   "list",
	Short: "List templates",
	Long:  "Query the Rewrite `/templates` endpoint with cursor options.",
	Example: `  rewrite template list
  rewrite template list --with-i18n`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		return commandtemplates.List(commandtemplates.ListOpts{
			RenderOptions: cliutil.ReadRenderOptions(cmd),
			Limit:         cliutil.ReadIntFlag(cmd, "limit"),
			Before:        cliutil.ReadStringFlag(cmd, "before"),
			After:         cliutil.ReadStringFlag(cmd, "after"),
			WithI18N:      cliutil.ReadBoolFlag(cmd, "with-i18n"),
		})
	},
}

func init() {
	templateListCmd.Flags().Int("limit", 20, "Maximum number of templates to return")
	templateListCmd.Flags().String("before", "", "Cursor for older templates")
	templateListCmd.Flags().String("after", "", "Cursor for newer templates")
	templateListCmd.Flags().Bool("with-i18n", false, "Include i18n values in the response")

	templateCmd.AddCommand(templateListCmd)
}
