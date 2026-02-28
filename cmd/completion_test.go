package cmd

import (
	"testing"

	"github.com/RewriteToday/cli/internal/clierr"
)

func TestCompletionArgs(t *testing.T) {
	if err := completionCmd.Args(completionCmd, []string{"zsh"}); err != nil {
		t.Fatalf("expected zsh to be valid: %v", err)
	}

	if err := completionCmd.Args(completionCmd, []string{}); err == nil {
		t.Fatal("expected missing shell argument error")
	}
}

func TestCompletionArgs_InvalidShellReturnsUsageError(t *testing.T) {
	err := completionCmd.Args(completionCmd, []string{"elvish"})
	if err == nil {
		t.Fatal("expected invalid shell error")
	}

	if got := clierr.ExitCode(err); got != int(clierr.CodeUsage) {
		t.Fatalf("expected usage exit code %d, got %d", clierr.CodeUsage, got)
	}
}
