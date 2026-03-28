package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	commandmessages "github.com/RewriteToday/cli/internal/commands/messages"
	"github.com/spf13/cobra"
)

var messageSendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a message through Rewrite",
	Long:  "Create a message with raw content or a stored template using the real `/messages` API.",
	Example: `  rewrite message send --to +5511999999999 --content "Hello from Rewrite"
  rewrite message send --to +5511999999999 --template-id 123 --variable name=Ana`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		tags, _ := cmd.Flags().GetStringArray("tag")
		variables, _ := cmd.Flags().GetStringArray("variable")

		return commandmessages.Send(commandmessages.SendOpts{
			RenderOptions:     cliutil.ReadRenderOptions(cmd),
			To:                cliutil.ReadStringFlag(cmd, "to"),
			Content:           cliutil.ReadStringFlag(cmd, "content"),
			TemplateID:        cliutil.ReadStringFlag(cmd, "template-id"),
			Variables:         variables,
			Tags:              tags,
			ScheduledAt:       cliutil.ReadStringFlag(cmd, "scheduled-at"),
			IdempotencyKey:    cliutil.ReadStringFlag(cmd, "idempotency-key"),
			SegmentationMax:   cliutil.ReadIntFlag(cmd, "segment-max"),
			SegmentationMode:  cliutil.ReadStringFlag(cmd, "segmentation-mode"),
			SegmentationSmart: cliutil.ReadBoolFlag(cmd, "smart-segmentation"),
		})
	},
}

func init() {
	messageSendCmd.Flags().String("to", "", "Destination phone number in E.164 format")
	messageSendCmd.Flags().String("content", "", "Raw SMS content")
	messageSendCmd.Flags().String("template-id", "", "Template ID to render instead of raw content")
	messageSendCmd.Flags().StringArray("variable", nil, "Template variable as name=value")
	messageSendCmd.Flags().StringArray("tag", nil, "Tag as name=value")
	messageSendCmd.Flags().String("scheduled-at", "", "Optional ISO-8601 schedule time")
	messageSendCmd.Flags().String("idempotency-key", "", "Optional idempotency key")
	messageSendCmd.Flags().Int("segment-max", 0, "Optional max number of segments")
	messageSendCmd.Flags().String("segmentation-mode", "", "Optional segmentation mode: fail, trim, or send")
	messageSendCmd.Flags().Bool("smart-segmentation", false, "Enable smart segmentation")

	messageCmd.AddCommand(messageSendCmd)
}
