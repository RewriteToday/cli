package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/RewriteToday/cli/internal/clierr"
)

const LocalListenRoute = "/events/listen"

var (
	localDispatchHTTPClient = &http.Client{Timeout: 5 * time.Second}
	localListenerStatePath  = filepath.Join(os.TempDir(), "rewrite-cli-listener.json")
)

type localListenerState struct {
	Port int `json:"port"`
	PID  int `json:"pid"`
}

type triggerEventMessage struct {
	Timestamp string `json:"timestamp"`
	EventType string `json:"event_type"`
	Payload   any    `json:"payload"`
}

func (c *Client) TriggerEvent(eventType EventType, data map[string]any) error {
	return DispatchEvent(eventType, data)
}

func DispatchEvent(eventType EventType, data map[string]any) error {
	state, err := readLocalListenerState()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return clierr.Errorf(clierr.CodeNotFound, "no active listener found, run 'rewrite listen' first")
		}

		return fmt.Errorf("failed to load local listener state: %w", err)
	}

	body, err := json.Marshal(triggerEventMessage{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		EventType: string(eventType),
		Payload:   data,
	})
	if err != nil {
		return fmt.Errorf("failed to encode event payload: %w", err)
	}

	endpoint := fmt.Sprintf("http://localhost:%d%s", state.Port, LocalListenRoute)
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to create local event request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := localDispatchHTTPClient.Do(req)
	if err != nil {
		var urlErr *url.Error
		if errors.As(err, &urlErr) {
			clearLocalListenerState(state)
			return clierr.Errorf(clierr.CodeNotFound, "no active listener responding on port %d, run 'rewrite listen' again", state.Port)
		}

		return clierr.Wrap(clierr.CodeNetwork, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return clierr.Errorf(clierr.CodeNetwork, "local listener rejected event with status %d", resp.StatusCode)
	}

	return nil
}

func RegisterLocalListener(port int) error {
	if !isValidListenerPort(port) {
		return clierr.Errorf(clierr.CodeUsage, "port must be between 1 and 65535")
	}

	data, err := json.Marshal(localListenerState{
		Port: port,
		PID:  os.Getpid(),
	})
	if err != nil {
		return fmt.Errorf("failed to encode local listener state: %w", err)
	}

	if err := os.WriteFile(localListenerStatePath, data, 0o600); err != nil {
		return fmt.Errorf("failed to persist local listener state: %w", err)
	}

	return nil
}

func UnregisterLocalListener(port int) error {
	state, err := readLocalListenerState()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return fmt.Errorf("failed to load local listener state: %w", err)
	}

	if state.Port != port || state.PID != os.Getpid() {
		return nil
	}

	if err := os.Remove(localListenerStatePath); err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("failed to clear local listener state: %w", err)
	}

	return nil
}

func readLocalListenerState() (localListenerState, error) {
	data, err := os.ReadFile(localListenerStatePath)
	if err != nil {
		return localListenerState{}, err
	}

	var state localListenerState
	if err := json.Unmarshal(data, &state); err != nil {
		return localListenerState{}, fmt.Errorf("failed to decode local listener state: %w", err)
	}

	if !isValidListenerPort(state.Port) {
		return localListenerState{}, fmt.Errorf("invalid local listener port %d", state.Port)
	}

	return state, nil
}

func clearLocalListenerState(state localListenerState) {
	current, err := readLocalListenerState()
	if err != nil || current != state {
		return
	}

	_ = os.Remove(localListenerStatePath)
}

func isValidListenerPort(port int) bool {
	return port >= 1 && port <= 65535
}
