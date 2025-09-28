// repo/app_repo.go
package repo

import (
	"database/sql"
	"time"

	"github.com/rasjonell/app-reviews/internal/domain"
)

type AppRepo struct {
	db *sql.DB
}

func NewAppRepo(db *sql.DB) *AppRepo {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS apps (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			app_id TEXT UNIQUE NOT NULL,
			name TEXT NOT NULL,
			enabled BOOLEAN NOT NULL DEFAULT TRUE,
			last_polled DATETIME
		);
	`)
	if err != nil {
		panic(err)
	}
	return &AppRepo{db: db}
}

func (a *AppRepo) AddAppIfNotExists(appID, name string) error {
	_, err := a.db.Exec(`
		INSERT OR IGNORE INTO apps (app_id, name, enabled) VALUES (?, ?, ?)
	`, appID, name, true)
	return err
}

func (a *AppRepo) GetAllApps() ([]*domain.App, error) {
	rows, err := a.db.Query("SELECT id, app_id, name, enabled, last_polled FROM apps")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var apps []*domain.App
	for rows.Next() {
		var app domain.App
		var polled *string
		err := rows.Scan(&app.ID, &app.AppID, &app.Name, &app.Enabled, &polled)
		if err != nil {
			return nil, err
		}
		if polled != nil {
			app.LastPolled = *polled
		}
		apps = append(apps, &app)
	}
	return apps, nil
}

func (a *AppRepo) SetLastPolled(appID string, t time.Time) error {
	_, err := a.db.Exec(`UPDATE apps SET last_polled = ? WHERE app_id = ?`, t, appID)
	return err
}

func (a *AppRepo) GetAppByAppID(appID string) (*domain.App, error) {
	row := a.db.QueryRow("SELECT id, app_id, name, enabled, last_polled FROM apps WHERE app_id = ?", appID)
	var app domain.App
	var lastPolled *string

	err := row.Scan(&app.ID, &app.AppID, &app.Name, &app.Enabled, &lastPolled)
	if err != nil {
		return nil, err
	}

	if lastPolled != nil {
		app.LastPolled = *lastPolled
	}

	return &app, nil
}
