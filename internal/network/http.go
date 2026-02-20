package network

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/RewriteToday/cli/internal/clierr"
)

const (
	DefaultLocalhostAddr = "localhost:8080"
	shutdownTimeout      = 5 * time.Second
)

func Serve(ctx context.Context, addr, route string, handler http.Handler) error {
	server, err := newServer(addr, route, handler)
	if err != nil {
		return err
	}

	listenErr := make(chan error, 1)
	go func() {
		listenErr <- server.ListenAndServe()
	}()

	select {
	case err := <-listenErr:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return clierr.Wrap(clierr.CodeNetwork, err)
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			return clierr.Wrap(clierr.CodeNetwork, err)
		}

		err := <-listenErr
		if errors.Is(err, http.ErrServerClosed) || err == nil {
			return nil
		}

		return clierr.Wrap(clierr.CodeNetwork, err)
	}
}

func newServer(addr, route string, handler http.Handler) (*http.Server, error) {
	if strings.TrimSpace(addr) == "" {
		return nil, fmt.Errorf("addr is required")
	}

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
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}, nil
}
