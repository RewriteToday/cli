package auth

import "testing"

func TestValidateAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{name: "valid", value: "rw_abc123.secret456"},
		{name: "trimmed", value: "  rw_abc123.secret456  "},
		{name: "missing prefix", value: "abc123.secret456", wantErr: true},
		{name: "missing secret", value: "rw_abc123.", wantErr: true},
		{name: "missing dot", value: "rw_abc123secret456", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, err := ValidateAPIKey(tt.value)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}

				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if value != "rw_abc123.secret456" {
				t.Fatalf("expected normalized API key, got %q", value)
			}
		})
	}
}
