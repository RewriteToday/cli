package cmd

import (
	"fmt"

	"github.com/rewritestudios/cli/internal/api"
	"github.com/rewritestudios/cli/internal/prompt"
	"github.com/spf13/cobra"
)

var triggerCmd = &cobra.Command{
	Use:   "trigger <event-type>",
	Short: "Trigger a test event",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		interactive, _ := cmd.Flags().GetBool("interactive")

		var eventTypeStr string

		if len(args) > 0 {
			eventTypeStr = args[0]
		} else if interactive {
			var err error
			eventTypeStr, err = prompt.SelectString("Select an event type", api.SupportedEventStrings())
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("event type required (or use -i for interactive mode)")
		}

		eventType, err := api.ValidateEventType(eventTypeStr)
		if err != nil {
			return err
		}

		data := api.MockData(eventType)

		if interactive {
			override, err := prompt.TriggerEventForm(eventTypeStr)
			if err != nil {
				return err
			}
			for k, v := range override {
				if v != "" {
					data[k] = v
				}
			}
		}

		client, err := api.New()
		if err != nil {
			return err
		}

		if err := client.TriggerEvent(eventType, data); err != nil {
			return err
		}

		fmt.Printf("Event '%s' triggered successfully.\n", eventType)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(triggerCmd)
}
