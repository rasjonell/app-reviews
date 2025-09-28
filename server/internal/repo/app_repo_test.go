package repo

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func newDB(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open in-memory sqlite: %v", err)
	}
	t.Cleanup(func() { db.Close() })
	return db
}

func TestAppRepo_AddAndGet(t *testing.T) {
	db := newDB(t)
	repo := NewAppRepo(db)

	if err := repo.AddAppIfNotExists("123", "Test App"); err != nil {
		t.Fatalf("AddAppIfNotExists failed: %v", err)
	}

	if err := repo.AddAppIfNotExists("123", "Test App"); err != nil {
		t.Fatalf("AddAppIfNotExists duplicate failed: %v", err)
	}

	apps, err := repo.GetAllApps()
	if err != nil {
		t.Fatalf("GetAllApps failed: %v", err)
	}
	if len(apps) != 1 {
		t.Fatalf("expected 1 app, got %d", len(apps))
	}
	if apps[0].AppID != "123" || apps[0].Name != "Test App" {
		t.Fatalf("unexpected app data: %+v", apps[0])
	}

	got, err := repo.GetAppByAppID("123")
	if err != nil {
		t.Fatalf("GetAppByAppID failed: %v", err)
	}
	if got.AppID != "123" {
		t.Fatalf("expected appId 123, got %s", got.AppID)
	}
}

func TestAppRepo_SetLastPolled(t *testing.T) {
	db := newDB(t)
	repo := NewAppRepo(db)

	if err := repo.AddAppIfNotExists("123", "Test App"); err != nil {
		t.Fatalf("AddAppIfNotExists failed: %v", err)
	}

	if err := repo.SetLastPolled("123", time.Now()); err != nil {
		t.Fatalf("SetLastPolled failed: %v", err)
	}

	got, err := repo.GetAppByAppID("123")
	if err != nil {
		t.Fatalf("GetAppByAppID failed: %v", err)
	}
	if got.LastPolled == "" {
		t.Fatalf("expected LastPolled to be set, got empty")
	}
}
