package network

import (
	"net/http"
	"testing"
)

func TestNewServer_ValidatesRequiredFields(t *testing.T) {
	handler := http.HandlerFunc(func(http.ResponseWriter, *http.Request) {})

	if _, err := newServer("", "/events", handler); err == nil {
		t.Fatal("expected addr validation error")
	}

	if _, err := newServer("localhost:8080", "", handler); err == nil {
		t.Fatal("expected route validation error")
	}

	if _, err := newServer("localhost:8080", "events", handler); err == nil {
		t.Fatal("expected route prefix validation error")
	}

	if _, err := newServer("localhost:8080", "/events", nil); err == nil {
		t.Fatal("expected handler validation error")
	}
}

func TestNewServer_ConfiguresTimeouts(t *testing.T) {
	handler := http.HandlerFunc(func(http.ResponseWriter, *http.Request) {})

	srv, err := newServer("localhost:8080", "/events", handler)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if srv.ReadHeaderTimeout <= 0 || srv.ReadTimeout <= 0 || srv.WriteTimeout <= 0 || srv.IdleTimeout <= 0 {
		t.Fatal("expected server timeouts to be configured")
	}
}
