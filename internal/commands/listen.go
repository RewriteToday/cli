package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/RewriteToday/cli/internal/api"
	cliutil "github.com/RewriteToday/cli/internal/cli"
	"github.com/RewriteToday/cli/internal/network"
	"github.com/RewriteToday/cli/internal/style"
)

type ListenOpts struct {
	cliutil.RenderOptions
	Port int
}

func Listen(opts ListenOpts) error {
	addr := fmt.Sprintf("localhost:%d", opts.Port)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := api.RegisterLocalListener(opts.Port); err != nil {
		return err
	}
	defer func() {
		_ = api.UnregisterLocalListener(opts.Port)
	}()

	fmt.Printf("Waiting for webhook events at http://%s%s (press Ctrl+C to stop)\n", addr, api.LocalListenRoute)

	return network.Serve(ctx, addr, api.LocalListenRoute, buildListenHandler(opts.Format, opts.NoColor))
}

func buildListenHandler(format string, noColor bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		defer r.Body.Close()

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read request body", http.StatusBadRequest)
			return
		}

		if format == "json" {
			fmt.Println(string(body))
			w.WriteHeader(http.StatusAccepted)
			return
		}

		var event style.EventMessage
		if err := json.Unmarshal(body, &event); err != nil {
			fmt.Println(string(body))
			w.WriteHeader(http.StatusAccepted)
			return
		}

		if event.Timestamp == "" {
			event.Timestamp = time.Now().Format(time.RFC3339)
		}

		if event.EventType == "" {
			event.EventType = "event.received"
		}

		if err := style.Print(event, format, noColor); err != nil {
			fmt.Println(string(body))
		}

		w.WriteHeader(http.StatusAccepted)
	})
}
