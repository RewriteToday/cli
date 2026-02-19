package render

import (
	"os"

	"golang.org/x/term"
)

func Hyperlink(text, url string, noColor bool) string {
	if !supportsHyperlinks() {
		return text + " -> " + url
	}

	return "\x1b]8;;" + url + "\x1b\\" + text + "\x1b]8;;\x1b\\"
}

func supportsHyperlinks() bool {
	return IS_COLOR_ENABLED &&
		os.Getenv("TERM") != "dumb" &&
		term.IsTerminal(int(os.Stdout.Fd()))
}
