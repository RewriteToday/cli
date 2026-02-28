package cmd

import "testing"

func TestShouldUseInteractive(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		interactive bool
		want        bool
	}{
		{
			name:        "flag forces interactive",
			args:        []string{"sms.created"},
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
			args:        []string{"sms.created"},
			interactive: false,
			want:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shouldUseInteractive(tt.args, tt.interactive); got != tt.want {
				t.Fatalf("expected %t, got %t", tt.want, got)
			}
		})
	}
}
