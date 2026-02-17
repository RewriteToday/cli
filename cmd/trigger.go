package cmd

import (
	"fmt"

	"github.com/rewritestudios/cli/internal/api"
	"github.com/rewritestudios/cli/internal/output"
	"github.com/rewritestudios/cli/internal/style"
	"github.com/spf13/cobra"
)

var triggerCmd = &cobra.Command{
	Use:   "trigger <event-type>",
	Short: "Trigger a test event",
	Args:  cobra.MaximumNArgs(1),
	RunE:  runTriggerCommand,
}

func runTriggerCommand(cmd *cobra.Command, args []string) error {
	interactive, _ := cmd.Flags().GetBool("interactive")
	format, _ := cmd.Flags().GetString("output")
	noColor, _ := cmd.Flags().GetBool("no-color")

	eventTypeStr, err := resolveTriggerEventType(args, interactive)
	if err != nil {
		return err
	}

	eventType, err := api.ValidateEventType(eventTypeStr)
	if err != nil {
		return err
	}

	data, err := buildTriggerPayload(eventType, eventTypeStr, interactive)
	if err != nil {
		return err
	}

	client, err := api.New()
	if err != nil {
		return err
	}

	if err := client.TriggerEvent(eventType, data); err != nil {
		return err
	}

	return output.Print(fmt.Sprintf("Event '%s' triggered successfully.", eventType), format, noColor)
}

func resolveTriggerEventType(args []string, interactive bool) (string, error) {
	var eventTypeStr string
	if len(args) > 0 {
		eventTypeStr = args[0]
	}

	if eventTypeStr == "" && interactive {
		selected, err := style.SelectString("Select an event type", api.SupportedEventStrings())
		if err != nil {
			return "", err
		}
		eventTypeStr = selected
	}

	if eventTypeStr == "" {
		return "", fmt.Errorf("event type required (or use -i for interactive mode)")
	}

	return eventTypeStr, nil
}

func buildTriggerPayload(eventType api.EventType, eventTypeStr string, interactive bool) (map[string]interface{}, error) {
	data := api.MockData(eventType)
	if !interactive {
		return data, nil
	}

	override, err := style.TriggerEventForm(eventTypeStr)
	if err != nil {
		return nil, err
	}

	for k, v := range override {
		if v != "" {
			data[k] = v
		}
	}

	return data, nil
}

func init() {
	rootCmd.AddCommand(triggerCmd)
}
