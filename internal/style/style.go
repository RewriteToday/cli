package style

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/RewriteToday/cli/internal/render"
	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/quick"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/charmbracelet/huh"
)

func InputString(title, placeholder string) (string, error) {
	var value string

	err := huh.NewInput().
		Title(title).
		Placeholder(placeholder).
		Value(&value).
		Run()
	if err != nil {
		return "", fmt.Errorf("input cancelled: %w", err)
	}

	return value, nil
}

func SelectString(title string, options []string) (string, error) {
	var value string

	opts := make([]huh.Option[string], len(options))
	for i, o := range options {
		opts[i] = huh.NewOption(o, o)
	}

	err := huh.NewSelect[string]().
		Title(title).
		Options(opts...).
		Value(&value).
		Run()
	if err != nil {
		return "", fmt.Errorf("selection cancelled: %w", err)
	}

	return value, nil
}

func Confirm(title string) (bool, error) {
	var value bool

	err := huh.NewConfirm().
		Title(title).
		Value(&value).
		Run()
	if err != nil {
		return false, fmt.Errorf("confirmation cancelled: %w", err)
	}

	return value, nil
}

func TriggerEventForm(eventType string) (map[string]any, error) {
	var to, from, body string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("To (phone number)").
				Placeholder("+5511999999999").
				Value(&to),
			huh.NewInput().
				Title("From (phone number)").
				Placeholder("+5511888888888").
				Value(&from),
			huh.NewInput().
				Title("Body (message content)").
				Placeholder("Hello from Rewrite!").
				Value(&body),
		),
	)

	if err := form.Run(); err != nil {
		return nil, fmt.Errorf("form cancelled: %w", err)
	}

	data := map[string]any{
		"to":   to,
		"from": from,
		"body": body,
	}

	return data, nil
}

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
		return printJSON(data, noColor)
	}

	return printText(data, noColor)
}

func PrintError(err error, format string) {
	if format == "json" {
		errData := map[string]string{"error": err.Error()}
		jsonErr := printJSONToWriter(os.Stderr, errData, false)
		if jsonErr == nil {
			return
		}
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}

	fmt.Fprintf(os.Stderr, "%s %s\n", render.Paint("Error:", render.Red, false), err.Error())
}

func printJSON(data any, noColor bool) error {
	return printJSONToWriter(os.Stdout, data, noColor)
}

func printJSONToWriter(w *os.File, data any, noColor bool) error {
	formatted, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	payload := string(formatted)
	if !strings.HasSuffix(payload, "\n") {
		payload += "\n"
	}

	if noColor || !render.IS_COLOR_ENABLED {
		_, err = fmt.Fprint(w, payload)
		return err
	}

	ensureVesperStyleRegistered()
	if err := quick.Highlight(w, payload, "json", "terminal256", vesperStyleName); err != nil {
		_, writeErr := fmt.Fprint(w, payload)
		if writeErr != nil {
			return writeErr
		}
		return nil
	}

	return nil
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
		return printJSON(data, noColor)
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

const vesperStyleName = "rewrite-json-minimal"

var registerVesperStyleOnce sync.Once

func ensureVesperStyleRegistered() {
	registerVesperStyleOnce.Do(func() {
		styles.Register(chroma.MustNewStyle(vesperStyleName, chroma.StyleEntries{
			chroma.Background:           "#0b0b0d bg:#0b0b0d",
			chroma.Text:                 "#d4d4d8",
			chroma.Punctuation:          "#71717a",
			chroma.Operator:             "#71717a",
			chroma.NameTag:              "#e4e4e7",
			chroma.LiteralString:        "#a1a1aa",
			chroma.LiteralStringDouble:  "#a1a1aa",
			chroma.LiteralNumber:        "#a1a1aa",
			chroma.LiteralNumberInteger: "#a1a1aa",
			chroma.Keyword:              "#a1a1aa",
			chroma.KeywordConstant:      "#a1a1aa",
		}))
	})
}
