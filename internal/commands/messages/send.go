package messages

import (
	"context"
	"strings"

	"github.com/RewriteToday/cli/internal/api"
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/clierr"
	"github.com/RewriteToday/cli/internal/style"
)

type SendOpts struct {
	cliutil.RenderOptions
	To                string
	Content           string
	TemplateID        string
	Variables         []string
	Tags              []string
	ScheduledAt       string
	IdempotencyKey    string
	SegmentationMax   int
	SegmentationMode  string
	SegmentationSmart bool
}

func Send(opts SendOpts) error {
	body, err := buildSendBody(opts)
	if err != nil {
		return err
	}

	client, err := api.New()
	if err != nil {
		return err
	}

	response, err := client.SendMessage(context.Background(), body, opts.IdempotencyKey)
	if err != nil {
		return err
	}

	return style.Print(response, opts.Format, opts.NoColor)
}

func buildSendBody(opts SendOpts) (map[string]any, error) {
	if strings.TrimSpace(opts.To) == "" {
		return nil, clierr.Errorf(clierr.CodeUsage, "missing required flag --to")
	}

	hasContent := strings.TrimSpace(opts.Content) != ""
	hasTemplate := strings.TrimSpace(opts.TemplateID) != ""

	if hasContent == hasTemplate {
		return nil, clierr.Errorf(
			clierr.CodeUsage,
			"set exactly one of --content or --template-id",
		)
	}

	tags, err := parseTags(opts.Tags)
	if err != nil {
		return nil, err
	}

	body := map[string]any{
		"to": opts.To,
	}

	if len(tags) > 0 {
		body["tags"] = tags
	}
	if strings.TrimSpace(opts.ScheduledAt) != "" {
		body["scheduledAt"] = strings.TrimSpace(opts.ScheduledAt)
	}

	segmentation := map[string]any{}
	if opts.SegmentationMax > 0 {
		segmentation["max"] = opts.SegmentationMax
	}
	if strings.TrimSpace(opts.SegmentationMode) != "" {
		segmentation["mode"] = strings.TrimSpace(opts.SegmentationMode)
	}
	if opts.SegmentationSmart {
		segmentation["smart"] = true
	}
	if len(segmentation) > 0 {
		body["segmentation"] = segmentation
	}

	if hasContent {
		body["content"] = strings.TrimSpace(opts.Content)
		return body, nil
	}

	variables, err := parseVariables(opts.Variables)
	if err != nil {
		return nil, err
	}

	body["templateId"] = strings.TrimSpace(opts.TemplateID)
	body["variables"] = variables

	return body, nil
}

func parseTags(values []string) ([]api.MessageTag, error) {
	if len(values) == 0 {
		return nil, nil
	}

	result := make([]api.MessageTag, 0, len(values))
	for _, value := range values {
		name, parsedValue, ok := strings.Cut(value, "=")
		if !ok || strings.TrimSpace(name) == "" || strings.TrimSpace(parsedValue) == "" {
			return nil, clierr.Errorf(
				clierr.CodeUsage,
				"invalid --tag value %q, expected name=value",
				value,
			)
		}

		result = append(result, api.MessageTag{
			Name:  strings.TrimSpace(name),
			Value: strings.TrimSpace(parsedValue),
		})
	}

	return result, nil
}

func parseVariables(values []string) (map[string]string, error) {
	result := map[string]string{}
	for _, value := range values {
		name, parsedValue, ok := strings.Cut(value, "=")
		if !ok || strings.TrimSpace(name) == "" || strings.TrimSpace(parsedValue) == "" {
			return nil, clierr.Errorf(
				clierr.CodeUsage,
				"invalid --variable value %q, expected name=value",
				value,
			)
		}

		result[strings.TrimSpace(name)] = strings.TrimSpace(parsedValue)
	}

	return result, nil
}
