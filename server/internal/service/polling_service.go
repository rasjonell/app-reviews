package service

import (
	"context"
	"log"
	"time"
)

func StartPollingJob(ctx context.Context, reviewSvc *ReviewService, appSvc *AppService) {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			apps, _ := appSvc.GetAllApps()
			for _, app := range apps {
				if app.Enabled {
					if err := reviewSvc.FetchAndStoreReviews(ctx, app.AppID); err != nil {
						log.Printf("[App %s] Polling failed: %v\n", app.Name, err)
					} else {
						log.Printf("[App %s] Polling done", app.Name)
					}
				}
			}
		case <-ctx.Done():
			return
		}
	}
}
