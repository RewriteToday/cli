package network

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/RewriteToday/cli/internal/render"
)

func OpenURL(url string, noColor bool) error {
	switch runtime.GOOS {
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	case "linux":
		return exec.Command("xdg-open", url).Start()
	default:
		_, err := fmt.Println(render.Hyperlink("Go to the docs (docs.rewritetoday.com)", "https://docs.rewritetoday.com", noColor))

		return err
	}
}
