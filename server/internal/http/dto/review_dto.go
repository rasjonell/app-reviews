package dto

import (
	"net/http"
	"time"

	"github.com/rasjonell/app-reviews/internal/domain"
)

type ReviewResponse struct {
	ID        int       `json:"id"`
	AppID     string    `json:"appId"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Rating    int       `json:"rating"`
	Timestamp time.Time `json:"timestamp"`
}

func NewReviewResponse(review *domain.Review) *ReviewResponse {
	return &ReviewResponse{
		ID:        review.ID,
		AppID:     review.AppID,
		Author:    review.Author,
		Title:     review.Title,
		Content:   review.Content,
		Rating:    review.Rating,
		Timestamp: review.Timestamp,
	}
}

type AppReviewsResponse struct {
	App     *AppResponse      `json:"app"`
	Reviews []*ReviewResponse `json:"reviews"`
}

func NewAppReviewsResponse(app *domain.App, reviews []*domain.Review) *AppReviewsResponse {
	reviewsDTO := make([]*ReviewResponse, 0, len(reviews))
	if len(reviews) == 0 {
		reviewsDTO = []*ReviewResponse{}
	}

	for _, r := range reviews {
		reviewsDTO = append(reviewsDTO, NewReviewResponse(r))
	}

	return &AppReviewsResponse{
		App:     NewAppResponse(app),
		Reviews: reviewsDTO,
	}
}

type AppReviewsRequst struct {
	AppID  string
	Limit  int
	Offset int
	Since  time.Time
}

func NewAppReviewsRequest(w http.ResponseWriter, r *http.Request) (*AppReviewsRequst, bool) {
	appID, ok := parseParams(w, r, "appId")
	if !ok {
		return nil, false
	}
	limit, offset := parsePagination(r)

	since := time.Now().Add(-(24 * 10) * time.Hour)
	sinceStr := r.URL.Query().Get("since")
	if sinceStr != "" {
		if parsed, err := time.Parse(time.RFC3339, sinceStr); err == nil {
			since = parsed
		}
	}

	return &AppReviewsRequst{
		AppID:  appID,
		Limit:  limit,
		Offset: offset,
		Since:  since,
	}, true
}

type PollRequst struct {
	AppID string
}

func NewPollRequest(w http.ResponseWriter, r *http.Request) (*PollRequst, bool) {
	appID, ok := parseParams(w, r, "appId")
	if !ok {
		return nil, false
	}

	return &PollRequst{
		AppID: appID,
	}, true

}
