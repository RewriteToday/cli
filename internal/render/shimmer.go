package render

import (
	"context"
	"fmt"
	"time"
)

const (
	reset = "\033[0m"

	base = "\033[38;5;245m"

	trail = "\033[37m"

	highlight = "\033[1;97m"

	clear = "\033[2K"
)

func Shimmer(ctx context.Context, text string) func() {
	cctx, cancel := context.WithCancel(ctx)

	go func() {
		r := []rune(text)
		n := len(r)
		pos := 0

		ticker := time.NewTicker(40 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-cctx.Done():
				fmt.Print("\r" + clear)
				return

			case <-ticker.C:
				out := ""

				for i := 0; i < n; i++ {
					d := i - pos
					if d < 0 {
						d = -d
					}

					switch {
					case d == 0:
						out += highlight + string(r[i]) + reset
					case d == 1:
						out += trail + string(r[i]) + reset
					default:
						out += base + string(r[i]) + reset
					}
				}

				fmt.Print("\r" + clear + out)

				pos++
				if pos > n+1 {
					pos = -1
				}
			}
		}
	}()

	return cancel
}
