package render

import "os"

func IsColorEnabled() bool {
	return os.Getenv("NO_COLOR") == ""
}

const (
	Reset = "\033[0m"
	Bold  = "\033[1m"

	Purple = "\033[38;5;141m"
	Gray   = "\033[90m"
	Blue   = "\033[34m"
	Red    = "\033[31m"
	Yellow = "\033[33m"

	VesperAccent = "\033[38;5;183m"
	VesperText   = "\033[38;5;252m"
	VesperMuted  = "\033[38;5;247m"
	VesperSubtle = "\033[38;5;242m"
	VesperBrand  = "\033[38;5;151m"
	VesperTeal   = "\033[38;5;116m"
	VesperCode   = "\033[38;5;150m"
	VesperGold   = "\033[38;5;229m"
)

func Paint(content, code string, noColor bool) string {
	if noColor || !IsColorEnabled() {
		return content
	}

	return code + content + Reset
}

func PaintAll(content string, noColor bool, codes ...string) string {
	if noColor || !IsColorEnabled() || len(codes) == 0 {
		return content
	}

	combined := ""
	for _, code := range codes {
		combined += code
	}

	return combined + content + Reset
}
