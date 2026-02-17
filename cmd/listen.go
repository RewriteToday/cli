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

var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Listen for real-time events",
	RunE: func(cmd *cobra.Command, args []string) error {
		format, _ := cmd.Flags().GetString("output")
		noColor, _ := cmd.Flags().GetBool("no-color")

		const route = "/events/listen"
		fmt.Printf("Waiting for webhook events at http://localhost:8080%s (press Ctrl+C to stop)\n", route)

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

			var event output.EventMessage
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

			if err := output.Print(event, "text", noColor); err != nil {
				fmt.Println(string(body))
			}

			w.WriteHeader(http.StatusAccepted)
		}))
	},
}

func init() {
	rootCmd.AddCommand(listenCmd)
}
