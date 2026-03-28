package templates

import (
	"context"
	"strings"

	"github.com/RewriteToday/cli/internal/api"
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/clierr"
	"github.com/RewriteToday/cli/internal/style"
)

type CreateOpts struct {
	cliutil.RenderOptions
	Name        string
	Content     string
	Description string
	Variables   []string
	Tags        []string
}

func Create(opts CreateOpts) error {
	body, err := buildTemplateBody(opts)
	if err != nil {
		return err
	}

	client, err := api.New()
	if err != nil {
		return err
	}

	response, err := client.CreateTemplate(context.Background(), body)
	if err != nil {
		return err
	}

	return style.Print(response, opts.Format, opts.NoColor)
}

func buildTemplateBody(opts CreateOpts) (map[string]any, error) {
	if strings.TrimSpace(opts.Name) == "" {
		return nil, clierr.Errorf(clierr.CodeUsage, "missing required flag --name")
	}
	if strings.TrimSpace(opts.Content) == "" {
		return nil, clierr.Errorf(clierr.CodeUsage, "missing required flag --content")
	}

	variables, err := parseTemplateVariables(opts.Variables)
	if err != nil {
		return nil, err
	}

	tags, err := parseTemplateTags(opts.Tags)
	if err != nil {
		return nil, err
	}

	body := map[string]any{
		"name":      strings.TrimSpace(opts.Name),
		"content":   strings.TrimSpace(opts.Content),
		"variables": variables,
	}
	if strings.TrimSpace(opts.Description) != "" {
		body["description"] = strings.TrimSpace(opts.Description)
	}
	if len(tags) > 0 {
		body["tags"] = tags
	}

	return body, nil
}

func parseTemplateVariables(values []string) ([]api.TemplateVariable, error) {
	if len(values) == 0 {
		return []api.TemplateVariable{}, nil
	}

	result := make([]api.TemplateVariable, 0, len(values))
	for _, value := range values {
		name, fallback, _ := strings.Cut(value, "=")
		if strings.TrimSpace(name) == "" {
			return nil, clierr.Errorf(
				clierr.CodeUsage,
				"invalid --variable value %q, expected name or name=fallback",
				value,
			)
		}

		item := api.TemplateVariable{
			Name: strings.TrimSpace(name),
		}
		if strings.TrimSpace(fallback) != "" {
			item.Fallback = strings.TrimSpace(fallback)
		}

		result = append(result, item)
	}

	return result, nil
}

func parseTemplateTags(values []string) ([]api.TemplateTag, error) {
	if len(values) == 0 {
		return nil, nil
	}

	result := make([]api.TemplateTag, 0, len(values))
	for _, value := range values {
		name, parsedValue, ok := strings.Cut(value, "=")
		if !ok || strings.TrimSpace(name) == "" || strings.TrimSpace(parsedValue) == "" {
			return nil, clierr.Errorf(
				clierr.CodeUsage,
				"invalid --tag value %q, expected name=value",
				value,
			)
		}

		result = append(result, api.TemplateTag{
			Name:  strings.TrimSpace(name),
			Value: strings.TrimSpace(parsedValue),
		})
	}

	return result, nil
}
