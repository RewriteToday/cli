package output

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rewritestudios/cli/internal/render"
)

type ProfileInfo struct {
	Name   string `json:"name"`
	APIKey string `json:"api_key"`
}

type ProfileListItem struct {
	Name   string `json:"name"`
	APIKey string `json:"api_key"`
	Active bool   `json:"active"`
}

type LogEntry struct {
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	EventType string `json:"event_type"`
	Status    string `json:"status"`
	Payload   any    `json:"payload,omitempty"`
}

type EventMessage struct {
	Timestamp string `json:"timestamp"`
	EventType string `json:"event_type"`
	Payload   any    `json:"payload"`
}

func Print(data any, format string, noColor bool) error {
	if format == "json" {
		return printJSON(data)
	}

	return printText(data, noColor)
}

func PrintError(err error, format string) {
	if format == "json" {
		errData := map[string]string{"error": err.Error()}
		data, _ := json.Marshal(errData)
		fmt.Fprintln(os.Stderr, string(data))
		return
	}

	fmt.Fprintf(os.Stderr, "%s %s\n", render.Paint("Error:", render.Red, false), err.Error())
}

func printJSON(data any) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func printText(data any, noColor bool) error {
	switch v := data.(type) {
	case ProfileInfo:
		fmt.Printf("%s %s\n", render.Paint("Profile:", render.Bold, noColor), v.Name)
		fmt.Printf("%s %s...\n", render.Paint("API Key:", render.Bold, noColor), maskKey(v.APIKey))

	case []ProfileListItem:
		if len(v) == 0 {
			fmt.Println("No profiles found. Run 'rewrite login' to create one.")
			return nil
		}
		for _, p := range v {
			marker := "  "
			if p.Active {
				marker = render.Paint("* ", render.Purple, noColor)
			}
			name := p.Name
			if p.Active {
				name = render.Paint(name, render.Bold, noColor)
			}
			fmt.Printf("%s%s %s\n", marker, name, render.Paint(maskKey(p.APIKey), render.Gray, noColor))
		}

	case EventMessage:
		fmt.Printf("%s %s %s\n",
			render.Paint(v.Timestamp, render.Gray, noColor),
			render.Paint(v.EventType, render.Purple, noColor),
			formatPayload(v.Payload))

	case LogEntry:
		fmt.Printf("%s %s %s %s\n",
			render.Paint(v.Timestamp, render.Gray, noColor),
			render.Paint(v.EventType, render.Purple, noColor),
			render.Paint(v.Status, render.Blue, noColor),
			v.ID)

	case []LogEntry:
		if len(v) == 0 {
			fmt.Println("No logs found.")
			return nil
		}
		for _, entry := range v {
			printText(entry, noColor)
		}

	case string:
		fmt.Println(v)

	default:
		return printJSON(data)
	}

	return nil
}

func maskKey(key string) string {
	if len(key) <= 12 {
		return key + "..."
	}
	return key[:12] + "..."
}

func formatPayload(payload any) string {
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Sprintf("%v", payload)
	}
	return string(data)
}
