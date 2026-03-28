package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	commandtemplates "github.com/RewriteToday/cli/internal/commands/templates"
	"github.com/spf13/cobra"
)

var templateGetCmd = &cobra.Command{
	Use:   "get [id-or-name]",
	Short: "Fetch one template",
	Long:  "Retrieve a template by numeric ID or unique name.",
	Args:  cobra.ExactArgs(1),
	Example: `  rewrite template get welcome
  rewrite template get 1234567890 --with-i18n`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return commandtemplates.Get(commandtemplates.GetOpts{
			RenderOptions: cliutil.ReadRenderOptions(cmd),
			Identifier:    args[0],
			WithI18N:      cliutil.ReadBoolFlag(cmd, "with-i18n"),
		})
	},
}

func init() {
	templateGetCmd.Flags().Bool("with-i18n", false, "Include i18n values in the response")
	templateCmd.AddCommand(templateGetCmd)
}
