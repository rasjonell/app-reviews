package repo

import (
	"database/sql"
	"time"

	"github.com/rasjonell/app-reviews/internal/domain"
)

type ReviewRepo struct {
	db *sql.DB
}

func NewReviewRepo(db *sql.DB) *ReviewRepo {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS reviews (
			id INTEGER PRIMARY KEY,
			app_id TEXT,
			author TEXT,
			title TEXT,
			content TEXT,
			rating INTEGER,
			timestamp DATETIME
		);
	`)
	if err != nil {
		panic(err)
	}
	return &ReviewRepo{db: db}
}

func (r *ReviewRepo) Store(review *domain.Review) error {
	_, err := r.db.Exec(`
		INSERT OR IGNORE INTO reviews (id, app_id, author, title, content, rating, timestamp)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`, review.ID, review.AppID, review.Author, review.Title, review.Content, review.Rating, review.Timestamp)
	return err
}

func (r *ReviewRepo) GetRecent(appID string, since time.Time, limit, offset int) ([]*domain.Review, error) {
	rows, err := r.db.Query(`
		SELECT id, app_id, author, title, content, rating, timestamp
		FROM reviews
		WHERE app_id = ? AND timestamp >= ?
		ORDER BY timestamp DESC
    LIMIT ? OFFSET ?
	`, appID, since, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []*domain.Review
	for rows.Next() {
		var r domain.Review
		err := rows.Scan(&r.ID, &r.AppID, &r.Author, &r.Title, &r.Content, &r.Rating, &r.Timestamp)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, &r)
	}
	return reviews, nil
}
