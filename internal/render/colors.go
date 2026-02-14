package render

import "os"

var IS_COLOR_ENABLED = os.Getenv("NO_COLOR") == ""

const (
	Reset = "\033[0m"
	Bold  = "\033[1m"

	Purple = "\033[38;5;141m"
	Gray   = "\033[90m"
	Blue   = "\033[34m"
	Red    = "\033[31m"
	Yellow = "\033[33m"
)

func Paint(content, code string, noColor bool) string {
	if noColor || !IS_COLOR_ENABLED {
		return content
	}

	return code + content + Reset
}
