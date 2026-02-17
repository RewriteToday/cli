package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/rewritestudios/cli/internal/network"
	"github.com/rewritestudios/cli/internal/output"
	"github.com/spf13/cobra"
)

var logsTailCmd = &cobra.Command{
	Use:   "tail",
	Short: "Receive logs via webhook in real-time",
	RunE: func(cmd *cobra.Command, args []string) error {
		format, _ := cmd.Flags().GetString("output")
		noColor, _ := cmd.Flags().GetBool("no-color")

		const route = "/logs/tail"
		fmt.Printf("Waiting for webhook logs at http://localhost:8080%s (press Ctrl+C to stop)\n", route)

		return network.Serve(route, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

			var entry output.LogEntry
			if err := json.Unmarshal(body, &entry); err != nil {
				fmt.Println(string(body))
				w.WriteHeader(http.StatusAccepted)
				return
			}

			if entry.Timestamp == "" {
				entry.Timestamp = time.Now().Format(time.RFC3339)
			}

			if err := output.Print(entry, "text", noColor); err != nil {
				fmt.Println(string(body))
			}

			w.WriteHeader(http.StatusAccepted)
		}))
	},
}

func init() {
	logsCmd.AddCommand(logsTailCmd)
}
