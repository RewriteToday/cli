package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/RewriteToday/cli/internal/network"
	"github.com/RewriteToday/cli/internal/style"
	"github.com/spf13/cobra"
)

var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Listen for real-time events",
	Example: `  rewrite listen
  rewrite listen --port 9090
  rewrite listen --output json`,
	RunE: runListenCommand,
}

func runListenCommand(cmd *cobra.Command, _ []string) error {
	format, _ := cmd.Flags().GetString("output")
	noColor, _ := cmd.Flags().GetBool("no-color")
	port, _ := cmd.Flags().GetInt("port")
	return listen(format, noColor, port)
}

func listen(format string, noColor bool, port int) error {
	const route = "/events/listen"
	addr := fmt.Sprintf("localhost:%d", port)
	fmt.Printf("Waiting for webhook events at http://%s%s (press Ctrl+C to stop)\n", addr, route)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	return network.Serve(ctx, addr, route, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	}))
}

func init() {
	listenCmd.Flags().Int("port", 8080, "Port to bind for local webhook listener")
	rootCmd.AddCommand(listenCmd)
}
