// handlers/handler.go
package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rasjonell/app-reviews/internal/http/dto"
	"github.com/rasjonell/app-reviews/internal/service"
)

type Handler struct {
	reviewSvc *service.ReviewService
	appSvc    *service.AppService
}

func NewHandler(svc *service.ReviewService, appSvc *service.AppService) *Handler {
	return &Handler{reviewSvc: svc, appSvc: appSvc}
}

func (h *Handler) AddApp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	req, err := dto.NewAppCreateRequest(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.appSvc.AddAppIfNotExists(req.AppID)
	if err != nil {
		http.Error(w, "Failed to add app: "+err.Error(), http.StatusInternalServerError)
		return
	}

	/*
	 * Immedietly trigger review fetching
	 * Because frontend will redirect to reviews page after creation
	 */
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if err := h.reviewSvc.FetchAndStoreReviews(ctx, req.AppID); err != nil {
			log.Printf("Initial fetch failed for %s: %v\n", req.AppID, err)
		}
	}()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func (h *Handler) GetApps(w http.ResponseWriter, r *http.Request) {
	apps, err := h.appSvc.GetAllApps()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto.NewAppsResponse(apps))
}

func (h *Handler) GetReviews(w http.ResponseWriter, r *http.Request) {
	req, ok := dto.NewAppReviewsRequest(w, r)
	if !ok {
		return
	}

	app, err := h.appSvc.GetAppByAppID(req.AppID)
	if err != nil {
		http.Error(w, "App not found", http.StatusNotFound)
		return
	}

	reviews, err := h.reviewSvc.GetRecentReviews(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto.NewAppReviewsResponse(app, reviews))
}

func (h *Handler) PollNow(w http.ResponseWriter, r *http.Request) {
	req, ok := dto.NewPollRequest(w, r)
	if !ok {
		return
	}

	err := h.reviewSvc.FetchAndStoreReviews(r.Context(), req.AppID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Poll failed: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Polled successfully"))
}
