package service

import (
	"database/sql"
	"io"
	"net/http"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rasjonell/app-reviews/internal/repo"
)

type roundTripFunc func(*http.Request) *http.Response

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r), nil
}

func newDB(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open in-memory sqlite: %v", err)
	}
	t.Cleanup(func() { db.Close() })
	return db
}

func TestAppService_AddAppIfNotExists_UsesAppStoreName(t *testing.T) {
	// mock appstore http call for app name lookup
	origTransport := http.DefaultClient.Transport
	t.Cleanup(func() { http.DefaultClient.Transport = origTransport })
	http.DefaultClient.Transport = roundTripFunc(func(req *http.Request) *http.Response {
		if strings.Contains(req.URL.String(), "/lookup?") {
			body := `{"results":[{"trackName":"Stubbed App Name"}]}`
			return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(body)), Header: make(http.Header)}
		}
		return &http.Response{StatusCode: 404, Body: io.NopCloser(strings.NewReader("{}")), Header: make(http.Header)}
	})

	db := newDB(t)
	appRepo := repo.NewAppRepo(db)
	appStore := NewAppStoreService()
	svc := NewAppService(appRepo, appStore)

	if err := svc.AddAppIfNotExists("999"); err != nil {
		t.Fatalf("AddAppIfNotExists failed: %v", err)
	}

	apps, err := svc.GetAllApps()
	if err != nil {
		t.Fatalf("GetAllApps failed: %v", err)
	}
	if len(apps) != 1 {
		t.Fatalf("expected 1 app, got %d", len(apps))
	}
	if apps[0].Name != "Stubbed App Name" {
		t.Fatalf("expected app name from stub, got %q", apps[0].Name)
	}
}
