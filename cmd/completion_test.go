package cmd

import "testing"

func TestCompletionArgs(t *testing.T) {
	if err := completionCmd.Args(completionCmd, []string{"zsh"}); err != nil {
		t.Fatalf("expected zsh to be valid: %v", err)
	}

	if err := completionCmd.Args(completionCmd, []string{}); err == nil {
		t.Fatal("expected missing shell argument error")
	}
}
