package network

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/RewriteToday/cli/internal/clierr"
)

const localhostAddr = "localhost:8080"

func Serve(route string, handler http.Handler) error {
	server, err := newServer(route, handler)
	if err != nil {
		return err
	}

	return clierr.Wrap(clierr.CodeNetwork, server.ListenAndServe())
}

func newServer(route string, handler http.Handler) (*http.Server, error) {
	if strings.TrimSpace(route) == "" {
		return nil, fmt.Errorf("route is required")
	}

	if !strings.HasPrefix(route, "/") {
		return nil, fmt.Errorf("route must start with '/'")
	}

	if handler == nil {
		return nil, fmt.Errorf("handler is required")
	}

	mux := http.NewServeMux()
	mux.Handle(route, handler)

	return &http.Server{
		Addr:    localhostAddr,
		Handler: mux,
	}, nil
}
