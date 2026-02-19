package network

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/RewriteToday/cli/internal/render"
)

func OpenURL(url string, noColor bool) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	case "linux":
		cmd = "xdg-open"
		args = []string{url}
	default:
		_, err := fmt.Println(render.Hyperlink("Go to the docs (docs.rewritetoday.com)", "https://docs.rewritetoday.com", noColor))

		return err
	}

	return exec.Command(cmd, args...).Start()
}
