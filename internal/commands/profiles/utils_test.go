package profiles

import (
	"testing"

	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/clierr"
)

func TestShouldUseInteractive(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		interactive bool
		want        bool
	}{
		{
			name:        "flag forces interactive",
			args:        []string{"team"},
			interactive: true,
			want:        true,
		},
		{
			name:        "missing arg enables interactive",
			args:        []string{},
			interactive: false,
			want:        true,
		},
		{
			name:        "provided arg keeps non-interactive",
			args:        []string{"team"},
			interactive: false,
			want:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cliutil.ShouldUseInteractive(tt.args, tt.interactive); got != tt.want {
				t.Fatalf("expected %t, got %t", tt.want, got)
			}
		})
	}
}

func TestResolveName_NonInteractiveRequiresArgument(t *testing.T) {
	_, err := resolveName([]string{}, false)
	if err == nil {
		t.Fatal("expected error")
	}

	if got := clierr.ExitCode(err); got != int(clierr.CodeUsage) {
		t.Fatalf("expected usage exit code %d, got %d", clierr.CodeUsage, got)
	}
}

func TestResolveName_UsesProvidedArgument(t *testing.T) {
	name, err := resolveName([]string{"team"}, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if name != "team" {
		t.Fatalf("expected team, got %s", name)
	}
}
