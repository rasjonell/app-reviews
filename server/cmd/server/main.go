package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/rasjonell/app-reviews/internal/db"
	"github.com/rasjonell/app-reviews/internal/http/handlers"
	"github.com/rasjonell/app-reviews/internal/http/middleware"
	"github.com/rasjonell/app-reviews/internal/repo"
	"github.com/rasjonell/app-reviews/internal/service"
)

const seedAppID = "595068606"

func main() {
	db := db.InitDB()
	defer db.Close()

	appRepo := repo.NewAppRepo(db)
	reviewRepo := repo.NewReviewRepo(db)

	appstoreSvc := service.NewAppStoreService()
	appSvc := service.NewAppService(appRepo, appstoreSvc)
	reviewSvc := service.NewReviewService(reviewRepo, appRepo, appstoreSvc)

	handler := handlers.NewHandler(reviewSvc, appSvc)

	// Seed
	appSvc.AddAppIfNotExists(seedAppID)
	reviewSvc.FetchAndStoreReviews(context.Background(), seedAppID)

	mux := http.NewServeMux()
	mux.HandleFunc("/apps", middleware.CorsMiddleware(handler.GetApps))
	mux.HandleFunc("/apps/new", middleware.CorsMiddleware(handler.AddApp))
	mux.HandleFunc("/apps/{appId}/reviews", middleware.CorsMiddleware(handler.GetReviews))
	mux.HandleFunc("/apps/{appId}/poll", middleware.CorsMiddleware(handler.PollNow))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	ctx, cancel := context.WithCancel(context.Background())
	go service.StartPollingJob(ctx, reviewSvc, appSvc)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	go func() {
		<-sigChan
		log.Println("Shutting down server...")
		cancel()
		server.Shutdown(ctx)
	}()

	log.Println("Server started on :8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
	}
}
