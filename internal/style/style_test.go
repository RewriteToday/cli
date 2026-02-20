package style

import "testing"

func TestSanitizeForJSON_ProfileInfoMasksAPIKey(t *testing.T) {
	in := ProfileInfo{Name: "dev", APIKey: "rw_live_1234567890abcdef"}

	out, ok := sanitizeForJSON(in).(ProfileInfoJSON)
	if !ok {
		t.Fatalf("expected ProfileInfoJSON type, got %T", sanitizeForJSON(in))
	}

	if out.Name != in.Name {
		t.Fatalf("expected name %q, got %q", in.Name, out.Name)
	}

	if out.APIKeyMasked == in.APIKey {
		t.Fatal("expected masked API key, got raw key")
	}

	if out.APIKeyMasked == "" {
		t.Fatal("expected non-empty masked key")
	}
}

func TestMaskKey(t *testing.T) {
	masked := MaskKey("rw_live_1234567890abcdef")
	if masked == "rw_live_1234567890abcdef" {
		t.Fatal("expected key to be masked")
	}
}
