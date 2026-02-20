package profiles

import (
	"testing"

	"github.com/RewriteToday/cli/internal/clierr"
)

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
