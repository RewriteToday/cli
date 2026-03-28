package messages

import (
	"testing"

	"github.com/RewriteToday/cli/internal/clierr"
)

func TestBuildSendBody(t *testing.T) {
	body, err := buildSendBody(SendOpts{
		To:              "+5511999999999",
		Content:         "Hello from Rewrite",
		Tags:            []string{"env=dev"},
		SegmentationMax: 2,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if body["to"] != "+5511999999999" {
		t.Fatalf("expected to field, got %#v", body["to"])
	}

	if body["content"] != "Hello from Rewrite" {
		t.Fatalf("expected content field, got %#v", body["content"])
	}

	segmentation, ok := body["segmentation"].(map[string]any)
	if !ok {
		t.Fatalf("expected segmentation map, got %T", body["segmentation"])
	}

	if segmentation["max"] != 2 {
		t.Fatalf("expected max 2, got %#v", segmentation["max"])
	}
}

func TestBuildSendBodyRequiresExactlyOnePayloadShape(t *testing.T) {
	_, err := buildSendBody(SendOpts{
		To:         "+5511999999999",
		Content:    "Hello",
		TemplateID: "123",
	})
	if clierr.CodeOf(err) != clierr.CodeUsage {
		t.Fatalf("expected usage error, got %v", err)
	}
}

func TestBuildSendBodyParsesTemplateVariables(t *testing.T) {
	body, err := buildSendBody(SendOpts{
		To:         "+5511999999999",
		TemplateID: "123",
		Variables:  []string{"name=Ana"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	variables, ok := body["variables"].(map[string]string)
	if !ok {
		t.Fatalf("expected variables map, got %T", body["variables"])
	}

	if variables["name"] != "Ana" {
		t.Fatalf("expected variable value, got %#v", variables["name"])
	}
}
