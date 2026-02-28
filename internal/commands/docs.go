package commands

import (
	"github.com/RewriteToday/cli/internal/config"
	"github.com/RewriteToday/cli/internal/network"
)

func Docs(noColor bool) error {
	return network.OpenURL(config.DocsURL, noColor)
}
