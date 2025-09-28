package service

import (
	"context"
	"time"

	"github.com/rasjonell/app-reviews/internal/domain"
	"github.com/rasjonell/app-reviews/internal/http/dto"
	"github.com/rasjonell/app-reviews/internal/repo"
)

type ReviewService struct {
	repo        *repo.ReviewRepo
	appRepo     *repo.AppRepo
	appStoreSvc *AppStoreService
}

func NewReviewService(repo *repo.ReviewRepo, appRepo *repo.AppRepo, appStoreSvc *AppStoreService) *ReviewService {
	return &ReviewService{repo: repo, appRepo: appRepo, appStoreSvc: appStoreSvc}
}

func (s *ReviewService) FetchAndStoreReviews(ctx context.Context, appID string) error {
	reviews, err := s.appStoreSvc.GetReviews(ctx, appID)
	if err != nil {
		return err
	}
	for _, review := range reviews {
		s.repo.Store(review)
	}

	s.appRepo.SetLastPolled(appID, time.Now())
	return nil
}

func (s *ReviewService) GetRecentReviews(req *dto.AppReviewsRequst) ([]*domain.Review, error) {
	return s.repo.GetRecent(req.AppID, req.Since, req.Limit, req.Offset)
}
