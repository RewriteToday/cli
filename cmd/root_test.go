package cmd

import (
	"testing"

	"github.com/RewriteToday/cli/internal/clierr"
)

func TestResolveOutputFormat(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{name: "default", args: []string{}, want: "text"},
		{name: "long flag", args: []string{"--output", "json"}, want: "json"},
		{name: "short flag", args: []string{"-o", "json"}, want: "json"},
		{name: "equals", args: []string{"--output=json"}, want: "json"},
		{name: "invalid falls back", args: []string{"--output=xml"}, want: "text"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ResolveOutputFormat(tt.args); got != tt.want {
				t.Fatalf("expected %q, got %q", tt.want, got)
			}
		})
	}
}

func TestNormalizeOutputFormat_InvalidReturnsUsageError(t *testing.T) {
	_, err := normalizeOutputFormat("xml")
	if err == nil {
		t.Fatal("expected an error")
	}

	if got := clierr.ExitCode(err); got != int(clierr.CodeUsage) {
		t.Fatalf("expected code %d, got %d", clierr.CodeUsage, got)
	}
}
