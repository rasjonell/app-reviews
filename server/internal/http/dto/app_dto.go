package dto

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/rasjonell/app-reviews/internal/domain"
)

type AppResponse struct {
	ID         int    `json:"id"`
	AppID      string `json:"appId"`
	Name       string `json:"name"`
	Enabled    bool   `json:"enabled"`
	LastPolled string `json:"lastPolled"`
}

func NewAppResponse(app *domain.App) *AppResponse {
	return &AppResponse{
		ID:         app.ID,
		AppID:      app.AppID,
		Name:       app.Name,
		Enabled:    app.Enabled,
		LastPolled: app.LastPolled,
	}
}

func NewAppsResponse(apps []*domain.App) []*AppResponse {
	if len(apps) == 0 {
		return []*AppResponse{}
	}

	dtos := make([]*AppResponse, 0, len(apps))
	for _, a := range apps {
		dtos = append(dtos, NewAppResponse(a))
	}

	return dtos
}

type AppCreateRequest struct {
	AppID string `json:"appId"`
}

func NewAppCreateRequest(reqBody io.ReadCloser) (*AppCreateRequest, error) {
	var body AppCreateRequest
	err := json.NewDecoder(reqBody).Decode(&body)
	if err != nil || body.AppID == "" {
		return nil, errors.New("invalid request payload")
	}

	return &body, nil
}
