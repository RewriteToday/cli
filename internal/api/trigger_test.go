package api

import (
	"context"
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/RewriteToday/cli/internal/clierr"
)

func TestRegisterAndUnregisterLocalListener(t *testing.T) {
	useTempListenerStatePath(t)

	if err := RegisterLocalListener(3000); err != nil {
		t.Fatalf("expected listener registration to succeed: %v", err)
	}

	state, err := readLocalListenerState()
	if err != nil {
		t.Fatalf("expected listener state to be readable: %v", err)
	}

	if state.Port != 3000 {
		t.Fatalf("expected port 3000, got %d", state.Port)
	}

	if state.PID != os.Getpid() {
		t.Fatalf("expected pid %d, got %d", os.Getpid(), state.PID)
	}

	if err := UnregisterLocalListener(3000); err != nil {
		t.Fatalf("expected listener unregister to succeed: %v", err)
	}

	if _, err := os.Stat(localListenerStatePath); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("expected listener state file to be removed, got %v", err)
	}
}

func TestDispatchEventSendsPayloadToActiveListener(t *testing.T) {
	useTempListenerStatePath(t)

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("expected test listener to start: %v", err)
	}
	defer listener.Close()

	messages := make(chan triggerEventMessage, 1)
	server := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != LocalListenRoute {
				http.NotFound(w, r)
				return
			}

			defer r.Body.Close()

			var message triggerEventMessage
			if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
				t.Errorf("expected payload to decode: %v", err)
				http.Error(w, "bad request", http.StatusBadRequest)
				return
			}

			messages <- message
			w.WriteHeader(http.StatusAccepted)
		}),
	}

	go func() {
		_ = server.Serve(listener)
	}()

	t.Cleanup(func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		_ = server.Shutdown(shutdownCtx)
	})

	port := listener.Addr().(*net.TCPAddr).Port
	if err := RegisterLocalListener(port); err != nil {
		t.Fatalf("expected listener registration to succeed: %v", err)
	}
	t.Cleanup(func() {
		_ = UnregisterLocalListener(port)
	})

	payload := map[string]any{"status": "failed"}
	if err := DispatchEvent(SMSFailed, payload); err != nil {
		t.Fatalf("expected dispatch to succeed: %v", err)
	}

	select {
	case message := <-messages:
		if message.EventType != string(SMSFailed) {
			t.Fatalf("expected event type %q, got %q", SMSFailed, message.EventType)
		}

		body, ok := message.Payload.(map[string]any)
		if !ok {
			t.Fatalf("expected payload map, got %T", message.Payload)
		}

		if body["status"] != "failed" {
			t.Fatalf("expected payload status to be %q, got %#v", "failed", body["status"])
		}

		if message.Timestamp == "" {
			t.Fatal("expected dispatch timestamp to be set")
		}
	case <-time.After(2 * time.Second):
		t.Fatal("timed out waiting for dispatched event")
	}
}

func TestDispatchEventWithoutActiveListenerFails(t *testing.T) {
	useTempListenerStatePath(t)

	err := DispatchEvent(SMSCreated, MockData(SMSCreated))
	if clierr.CodeOf(err) != clierr.CodeNotFound {
		t.Fatalf("expected not_found error, got %v", err)
	}
}

func useTempListenerStatePath(t *testing.T) {
	t.Helper()

	original := localListenerStatePath
	localListenerStatePath = filepath.Join(t.TempDir(), "listener-state.json")

	t.Cleanup(func() {
		localListenerStatePath = original
	})
}
