package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/rasjonell/app-reviews/internal/domain"
)

type AppStoreService struct{}

func NewAppStoreService() *AppStoreService { return &AppStoreService{} }

func (s *AppStoreService) GetAppName(appID string) (string, error) {
	url := fmt.Sprintf("https://itunes.apple.com/lookup?id=%s", appID)
	req, _ := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var data map[string]any
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	results, ok := data["results"].([]any)
	if !ok || len(results) == 0 {
		return "", fmt.Errorf("no results found")
	}

	appInfo, ok := results[0].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("invalid response structure")
	}

	name, ok := appInfo["trackName"].(string)
	if !ok {
		return "", fmt.Errorf("app name not found")
	}

	return name, nil
}

func (s *AppStoreService) GetReviews(ctx context.Context, appID string) ([]*domain.Review, error) {
	url := fmt.Sprintf("https://itunes.apple.com/us/rss/customerreviews/id=%s/sortBy=mostRecent/page=1/json", appID)
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	feed, ok := data["feed"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response structure")
	}

	entries, ok := feed["entry"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("no reviews found")
	}

	reviews := make([]*domain.Review, 0, len(entries))
	for _, entry := range entries {
		e, _ := entry.(map[string]interface{})
		idStr, _ := e["id"].(map[string]interface{})["label"].(string)
		author, _ := e["author"].(map[string]interface{})["name"].(map[string]interface{})["label"].(string)
		title, _ := e["title"].(map[string]interface{})["label"].(string)
		content, _ := e["content"].(map[string]interface{})["label"].(string)
		ratingStr, _ := e["im:rating"].(map[string]interface{})["label"].(string)
		updatedStr, _ := e["updated"].(map[string]interface{})["label"].(string)

		reviewID := 0
		fmt.Sscanf(idStr, "%d", &reviewID)

		timestamp, _ := time.Parse(time.RFC3339, updatedStr)
		rating := 0
		fmt.Sscanf(ratingStr, "%d", &rating)

		reviews = append(reviews, &domain.Review{
			ID:        reviewID,
			AppID:     appID,
			Author:    author,
			Title:     title,
			Content:   content,
			Rating:    rating,
			Timestamp: timestamp,
		})
	}

	return reviews, nil
}
