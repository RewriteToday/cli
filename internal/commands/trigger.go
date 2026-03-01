package commands

import (
	"fmt"

	"github.com/RewriteToday/cli/internal/api"
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/clierr"
	"github.com/RewriteToday/cli/internal/style"
)

type TriggerOpts struct {
	cliutil.InteractiveRenderOptions
	Args []string
}

func Trigger(opts TriggerOpts) error {
	interactive := cliutil.ShouldUseInteractive(opts.Args, opts.Interactive)

	eventTypeStr, err := resolveTriggerEventType(opts.Args, interactive)
	if err != nil {
		return err
	}

	eventType, err := api.ValidateEventType(eventTypeStr)
	if err != nil {
		return err
	}

	data, err := buildTriggerPayload(eventType, interactive)
	if err != nil {
		return err
	}

	if err := api.DispatchEvent(eventType, data); err != nil {
		return err
	}

	return style.Print(fmt.Sprintf("Event '%s' triggered successfully.", eventType), opts.Format, opts.NoColor)
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
		return "", clierr.Errorf(clierr.CodeUsage, "event type required (or use -i for interactive mode)")
	}

	return eventTypeStr, nil
}

func buildTriggerPayload(eventType api.EventType, interactive bool) (map[string]any, error) {
	data := api.MockData(eventType)
	if !interactive {
		return data, nil
	}

	override, err := style.TriggerEventForm()
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
