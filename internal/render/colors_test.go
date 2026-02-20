package render

import "testing"

func TestIsColorEnabled_RespectsNoColor(t *testing.T) {
	t.Setenv("NO_COLOR", "1")
	if IsColorEnabled() {
		t.Fatal("expected colors to be disabled when NO_COLOR is set")
	}
}

func TestIsColorEnabled_DefaultEnabled(t *testing.T) {
	t.Setenv("NO_COLOR", "")
	if !IsColorEnabled() {
		t.Fatal("expected colors to be enabled when NO_COLOR is empty")
	}
}
