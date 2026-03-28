package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	commandwebhooks "github.com/RewriteToday/cli/internal/commands/webhooks"
	"github.com/spf13/cobra"
)

var webhookCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a webhook",
	Long:  "Register a webhook endpoint with one or more subscribed events.",
	Example: `  rewrite webhook create --endpoint https://example.com/webhooks --event message.sent
  rewrite webhook create --name billing --endpoint https://example.com/hooks --event '*'`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		events, _ := cmd.Flags().GetStringArray("event")

		return commandwebhooks.Create(commandwebhooks.CreateOpts{
			RenderOptions: cliutil.ReadRenderOptions(cmd),
			Name:          cliutil.ReadStringFlag(cmd, "name"),
			Endpoint:      cliutil.ReadStringFlag(cmd, "endpoint"),
			Secret:        cliutil.ReadStringFlag(cmd, "secret"),
			Events:        events,
		})
	},
}

func init() {
	webhookCreateCmd.Flags().String("name", "", "Optional webhook name")
	webhookCreateCmd.Flags().String("endpoint", "", "Webhook endpoint URL")
	webhookCreateCmd.Flags().String("secret", "", "Optional webhook secret")
	webhookCreateCmd.Flags().StringArray("event", nil, "Subscribed event; repeat for multiple events")

	webhookCmd.AddCommand(webhookCreateCmd)
}
