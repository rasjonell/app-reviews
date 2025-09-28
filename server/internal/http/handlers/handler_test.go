package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rasjonell/app-reviews/internal/http/dto"
	"github.com/rasjonell/app-reviews/internal/repo"
	"github.com/rasjonell/app-reviews/internal/service"
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

func TestHandler_GetApps(t *testing.T) {
	db := newDB(t)
	appRepo := repo.NewAppRepo(db)
	if err := appRepo.AddAppIfNotExists("a1", "App One"); err != nil {
		t.Fatal(err)
	}
	if err := appRepo.AddAppIfNotExists("a2", "App Two"); err != nil {
		t.Fatal(err)
	}

	appSvc := service.NewAppService(appRepo, service.NewAppStoreService())
	h := NewHandler(nil, appSvc)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/apps", nil)
	h.GetApps(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}

	var resp []*dto.AppResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if len(resp) != 2 {
		t.Fatalf("expected 2 apps, got %d", len(resp))
	}
}
