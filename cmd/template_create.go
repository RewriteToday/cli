package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	commandtemplates "github.com/RewriteToday/cli/internal/commands/templates"
	"github.com/spf13/cobra"
)

var templateCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a template",
	Long:  "Create a new SMS template through the Rewrite `/templates` API.",
	Example: `  rewrite template create --name welcome --content "Hello {{name}}"
  rewrite template create --name welcome --content "Hello {{name}}" --variable name=friend`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		tags, _ := cmd.Flags().GetStringArray("tag")
		variables, _ := cmd.Flags().GetStringArray("variable")

		return commandtemplates.Create(commandtemplates.CreateOpts{
			RenderOptions: cliutil.ReadRenderOptions(cmd),
			Name:          cliutil.ReadStringFlag(cmd, "name"),
			Content:       cliutil.ReadStringFlag(cmd, "content"),
			Description:   cliutil.ReadStringFlag(cmd, "description"),
			Variables:     variables,
			Tags:          tags,
		})
	},
}

func init() {
	templateCreateCmd.Flags().String("name", "", "Template name")
	templateCreateCmd.Flags().String("content", "", "Template content")
	templateCreateCmd.Flags().String("description", "", "Optional template description")
	templateCreateCmd.Flags().StringArray("variable", nil, "Variable as name or name=fallback")
	templateCreateCmd.Flags().StringArray("tag", nil, "Tag as name=value")

	templateCmd.AddCommand(templateCreateCmd)
}
