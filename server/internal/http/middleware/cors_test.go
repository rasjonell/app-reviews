package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCorsMiddleware_GET(t *testing.T) {
	next := func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("ok")) }
	handler := CorsMiddleware(next)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
	if rr.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Fatalf("missing CORS headers")
	}
}

func TestCorsMiddleware_OPTIONS(t *testing.T) {
	next := func(w http.ResponseWriter, r *http.Request) { t.Fatalf("next should not be called for OPTIONS") }
	handler := CorsMiddleware(next)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodOptions, "/", nil)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200 for OPTIONS, got %d", rr.Code)
	}
}
