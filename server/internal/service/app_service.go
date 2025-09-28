package service

import (
	"github.com/rasjonell/app-reviews/internal/domain"
	"github.com/rasjonell/app-reviews/internal/repo"
)

type AppService struct {
	repo        *repo.AppRepo
	appStoreSvc *AppStoreService
}

func NewAppService(repo *repo.AppRepo, appStoreSvc *AppStoreService) *AppService {
	return &AppService{repo: repo, appStoreSvc: appStoreSvc}
}

func (s *AppService) GetAllApps() ([]*domain.App, error) {
	return s.repo.GetAllApps()
}

func (s *AppService) AddAppIfNotExists(appID string) error {
	name, err := s.appStoreSvc.GetAppName(appID)
	if err != nil {
		return err
	}
	return s.repo.AddAppIfNotExists(appID, name)
}

func (s *AppService) GetAppByAppID(appID string) (*domain.App, error) {
	return s.repo.GetAppByAppID(appID)
}
