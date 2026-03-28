package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	commandtemplates "github.com/RewriteToday/cli/internal/commands/templates"
	"github.com/spf13/cobra"
)

var templateDeleteCmd = &cobra.Command{
	Use:     "delete [id]",
	Short:   "Delete a template",
	Long:    "Delete a template by ID through the Rewrite `/templates/:id` endpoint.",
	Aliases: []string{"remove", "rm"},
	Args:    cobra.ExactArgs(1),
	Example: `  rewrite template delete 1234567890`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return commandtemplates.Delete(commandtemplates.DeleteOpts{
			RenderOptions: cliutil.ReadRenderOptions(cmd),
			ID:            args[0],
		})
	},
}

func init() {
	templateCmd.AddCommand(templateDeleteCmd)
}
