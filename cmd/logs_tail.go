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

var logsTailCmd = &cobra.Command{
	Use:   "tail",
	Short: "Receive logs via webhook in real-time",
	Example: `  rewrite logs tail
  rewrite logs tail --port 9090
  rewrite logs tail --output json`,
	RunE: runLogsTailCommand,
}

func runLogsTailCommand(cmd *cobra.Command, _ []string) error {
	format, _ := cmd.Flags().GetString("output")
	noColor, _ := cmd.Flags().GetBool("no-color")
	port, _ := cmd.Flags().GetInt("port")
	return tailLogs(format, noColor, port)
}

func tailLogs(format string, noColor bool, port int) error {
	const route = "/logs/tail"
	addr := fmt.Sprintf("localhost:%d", port)
	fmt.Printf("Waiting for webhook logs at http://%s%s (press Ctrl+C to stop)\n", addr, route)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	return network.Serve(ctx, addr, route, buildLogsTailHandler(format, noColor))
}

func buildLogsTailHandler(format string, noColor bool) http.Handler {
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

		var entry style.LogEntry
		if err := json.Unmarshal(body, &entry); err != nil {
			fmt.Println(string(body))
			w.WriteHeader(http.StatusAccepted)
			return
		}

		if entry.Timestamp == "" {
			entry.Timestamp = time.Now().Format(time.RFC3339)
		}

		if err := style.Print(entry, format, noColor); err != nil {
			fmt.Println(string(body))
		}

		w.WriteHeader(http.StatusAccepted)
	})
}

func init() {
	logsTailCmd.Flags().Int("port", 8080, "Port to bind for local webhook log listener")
	logsCmd.AddCommand(logsTailCmd)
}
