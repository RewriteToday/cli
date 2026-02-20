package profiles

import (
	"testing"

	"github.com/RewriteToday/cli/internal/style"
)

func TestCreateProfileListJSONItems_MasksAPIKey(t *testing.T) {
	items := createProfileListJSONItems([]style.ProfileListItem{
		{Name: "dev", APIKey: "rw_live_1234567890abcdef"},
	}, "dev")

	if len(items) != 1 {
		t.Fatalf("expected one item, got %d", len(items))
	}

	if items[0].APIKeyMasked == "rw_live_1234567890abcdef" {
		t.Fatal("expected API key to be masked")
	}

	if !items[0].Active {
		t.Fatal("expected active profile to be flagged")
	}
}
