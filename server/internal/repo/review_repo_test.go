package repo

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rasjonell/app-reviews/internal/domain"
)

func newTestDBReviews(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open in-memory sqlite: %v", err)
	}
	t.Cleanup(func() { db.Close() })
	return db
}

func TestReviewRepo_StoreAndGetRecent(t *testing.T) {
	db := newTestDBReviews(t)
	rrepo := NewReviewRepo(db)

	now := time.Now()
	older := now.Add(-2 * time.Hour)

	reviews := []*domain.Review{
		{ID: 1, AppID: "appA", Author: "A", Title: "t1", Content: "c1", Rating: 4, Timestamp: now},
		{ID: 2, AppID: "appA", Author: "B", Title: "t2", Content: "c2", Rating: 5, Timestamp: older},
	}
	for _, rv := range reviews {
		if err := rrepo.Store(rv); err != nil {
			t.Fatalf("Store failed: %v", err)
		}
	}

	out, err := rrepo.GetRecent("appA", now.Add(-3*time.Hour), 10, 0)
	if err != nil {
		t.Fatalf("GetRecent failed: %v", err)
	}
	if len(out) != 2 {
		t.Fatalf("expected 2 reviews, got %d", len(out))
	}
	if out[0].ID != 1 || out[1].ID != 2 {
		t.Fatalf("expected order [1,2], got [%d,%d]", out[0].ID, out[1].ID)
	}

	// Pagination
	out2, err := rrepo.GetRecent("appA", now.Add(-3*time.Hour), 1, 1)
	if err != nil {
		t.Fatalf("GetRecent with pagination failed: %v", err)
	}
	if len(out2) != 1 || out2[0].ID != 2 {
		t.Fatalf("expected second review with ID 2, got %+v", out2)
	}
}
