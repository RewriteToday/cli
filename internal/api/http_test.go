package api

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/RewriteToday/cli/internal/clierr"
)

func TestClientEndpointPreservesBasePath(t *testing.T) {
	client := &Client{
		BaseURL: "https://api.rewritetoday.com/v1",
	}

	query := url.Values{
		"limit": []string{"10"},
	}

	endpoint, err := client.endpoint("/webhooks", query)
	if err != nil {
		t.Fatalf("expected endpoint to build successfully: %v", err)
	}

	expected := "https://api.rewritetoday.com/v1/webhooks?limit=10"
	if endpoint != expected {
		t.Fatalf("expected endpoint %q, got %q", expected, endpoint)
	}
}

func TestDecodeResponseReturnsErrorForEmptyNon2xxBody(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusNotFound,
		Body:       io.NopCloser(strings.NewReader(" \n\t ")),
	}

	_, err := decodeResponse(resp, nil)
	if err == nil {
		t.Fatal("expected empty non-2xx response to return an error")
	}

	if clierr.CodeOf(err) != clierr.CodeNotFound {
		t.Fatalf("expected error code %s, got %s", clierr.CodeNotFound, clierr.CodeOf(err))
	}
}

func TestDecodeResponseAllowsEmptySuccessfulBody(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusNoContent,
		Body:       io.NopCloser(strings.NewReader("")),
	}

	cursor, err := decodeResponse(resp, nil)
	if err != nil {
		t.Fatalf("expected empty successful response to return no error: %v", err)
	}

	if cursor != nil {
		t.Fatalf("expected no cursor, got %#v", cursor)
	}
}
