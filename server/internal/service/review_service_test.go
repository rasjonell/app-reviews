package service

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rasjonell/app-reviews/internal/repo"
)

func newTestDBSvc(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open in-memory sqlite: %v", err)
	}
	t.Cleanup(func() { db.Close() })
	return db
}

func TestReviewService_FetchAndStoreReviews(t *testing.T) {
	// mock appstore http call for reviews feed
	orig := http.DefaultClient.Transport
	t.Cleanup(func() { http.DefaultClient.Transport = orig })
	http.DefaultClient.Transport = roundTripFunc(func(req *http.Request) *http.Response {
		url := req.URL.String()
		switch {
		case strings.Contains(url, "/customerreviews/"):
			t1 := time.Now().Add(-1 * time.Hour).Format(time.RFC3339)
			t2 := time.Now().Add(-2 * time.Hour).Format(time.RFC3339)
			body := fmt.Sprintf(`{
        "feed": {"entry": [
            {"id": {"label": "1"},
             "author": {"name": {"label": "Alice"}},
             "title": {"label": "Great"},
             "content": {"label": "Works well"},
             "im:rating": {"label": "5"},
             "updated": {"label": %q}
            },
            {"id": {"label": "2"},
             "author": {"name": {"label": "Bob"}},
             "title": {"label": "Okay"},
             "content": {"label": "It's fine"},
             "im:rating": {"label": "4"},
             "updated": {"label": %q}
            }
        ]}
        }`, t1, t2)
			return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(body)), Header: make(http.Header)}
		default:
			return &http.Response{StatusCode: 404, Body: io.NopCloser(strings.NewReader("{}")), Header: make(http.Header)}
		}
	})

	db := newTestDBSvc(t)
	appRepo := repo.NewAppRepo(db)
	reviewRepo := repo.NewReviewRepo(db)
	appStore := NewAppStoreService()
	svc := NewReviewService(reviewRepo, appRepo, appStore)

	if err := appRepo.AddAppIfNotExists("appX", "App X"); err != nil {
		t.Fatalf("seed app failed: %v", err)
	}

	if err := svc.FetchAndStoreReviews(context.Background(), "appX"); err != nil {
		t.Fatalf("FetchAndStoreReviews returned error: %v", err)
	}

	out, err := reviewRepo.GetRecent("appX", time.Now().Add(-24*time.Hour), 10, 0)
	if err != nil {
		t.Fatalf("GetRecent failed: %v", err)
	}
	if len(out) != 2 {
		t.Fatalf("expected 2 reviews, got %d", len(out))
	}

	app, err := appRepo.GetAppByAppID("appX")
	if err != nil {
		t.Fatalf("GetAppByAppID failed: %v", err)
	}
	if app.LastPolled == "" {
		t.Fatalf("expected LastPolled to be set")
	}
}
