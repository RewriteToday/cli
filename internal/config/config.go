package config

import (
	"os"
	"strings"
)

const (
	DefaultAPIBaseURL = "https://api.rewritetoday.com/v1"
	DocsURL           = "https://docs.rewritetoday.com"
	KeyringService    = "rewrite-cli"
	ProfilesKey       = "rewrite_profiles"
	ActiveKey         = "rewrite_active_profile"

	APIBaseURLEnvVar = "REWRITE_API_BASE_URL"
)

func APIBaseURL() string {
	if value := strings.TrimSpace(os.Getenv(APIBaseURLEnvVar)); value != "" {
		return strings.TrimRight(value, "/")
	}

	return DefaultAPIBaseURL
}
