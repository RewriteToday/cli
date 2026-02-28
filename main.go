package main

import (
	"os"

	"github.com/RewriteToday/cli/cmd"
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/clierr"
	"github.com/RewriteToday/cli/internal/style"
)

func main() {
	if err := cmd.Execute(); err != nil {
		format := cliutil.ResolveOutputFormat(os.Args[1:])

		style.PrintError(err, format)

		os.Exit(clierr.ExitCode(err))
	}
}
