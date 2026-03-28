package auth

import (
	"strings"

	"github.com/RewriteToday/cli/internal/clierr"
)

func IsAPIKey(value string) bool {
	apiKey := strings.TrimSpace(value)
	if !strings.HasPrefix(apiKey, "rw_") {
		return false
	}

	dot := strings.IndexByte(apiKey, '.')
	return dot > len("rw_") && dot < len(apiKey)-1
}

func ValidateAPIKey(value string) (string, error) {
	apiKey := strings.TrimSpace(value)
	if !IsAPIKey(apiKey) {
		return "", clierr.Errorf(clierr.CodeUsage, "invalid Rewrite API key format")
	}

	return apiKey, nil
}
