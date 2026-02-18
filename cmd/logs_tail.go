package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/RewriteToday/cli/internal/network"
	"github.com/RewriteToday/cli/internal/style"
	"github.com/spf13/cobra"
)

var logsTailCmd = &cobra.Command{
	Use:   "tail",
	Short: "Receive logs via webhook in real-time",
	RunE:  runLogsTailCommand,
}

func runLogsTailCommand(cmd *cobra.Command, _ []string) error {
	format, _ := cmd.Flags().GetString("output")
	noColor, _ := cmd.Flags().GetBool("no-color")
	return tailLogs(format, noColor)
}

func tailLogs(format string, noColor bool) error {
	const route = "/logs/tail"
	fmt.Printf("Waiting for webhook logs at http://localhost:8080%s (press Ctrl+C to stop)\n", route)
	return network.Serve(route, buildLogsTailHandler(format, noColor))
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
	logsCmd.AddCommand(logsTailCmd)
}
