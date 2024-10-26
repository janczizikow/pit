package server

import (
	"net/http"

	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/janczizikow/pit/internal/handlers"
	"github.com/janczizikow/pit/internal/http/middleware"
	"github.com/janczizikow/pit/internal/repository"
)

func Router(s *Server) http.Handler {
	mux := http.NewServeMux()
	sentryHandler := sentryhttp.New(sentryhttp.Options{
		Repanic: true,
	})

	repo := repository.New(s.db)
	seasonsHandler := handlers.NewSeasonsHandler(repo.Seasons)
	seasonSubmissionsHandler := handlers.NewSeasonSubmissionsHandler(repo.SeasonSubmissions)

	mux.HandleFunc("GET /api/v1/seasons", sentryHandler.HandleFunc(seasonsHandler.ListSeasons))
	mux.Handle("GET /api/v1/seasons/current", sentryHandler.HandleFunc(seasonsHandler.Current))
	mux.HandleFunc("GET /api/v1/seasons/{id}/submissions", sentryHandler.HandleFunc(seasonSubmissionsHandler.ListSubmissions))
	mux.HandleFunc("POST /api/v1/seasons/{id}/submissions", sentryHandler.HandleFunc(seasonSubmissionsHandler.CreateSubmission))
	mux.HandleFunc("GET /api/v1/seasons/{id}/statistics", sentryHandler.HandleFunc(seasonsHandler.GetStatistics))

	return middleware.Recover(middleware.Compression(mux))
}
