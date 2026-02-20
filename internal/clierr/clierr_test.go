package clierr

import (
	"context"
	"errors"
	"testing"
)

func TestExitCode_WithWrappedCode(t *testing.T) {
	err := Wrap(CodeUsage, errors.New("bad args"))

	if got := ExitCode(err); got != int(CodeUsage) {
		t.Fatalf("expected %d, got %d", CodeUsage, got)
	}
}

func TestExitCode_WithNetworkDeadline(t *testing.T) {
	err := context.DeadlineExceeded

	if got := ExitCode(err); got != int(CodeNetwork) {
		t.Fatalf("expected %d, got %d", CodeNetwork, got)
	}
}

func TestExitCode_WithUnknownError(t *testing.T) {
	err := errors.New("boom")

	if got := ExitCode(err); got != int(CodeInternal) {
		t.Fatalf("expected %d, got %d", CodeInternal, got)
	}
}
