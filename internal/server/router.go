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

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://diablo4pit.web.app/", http.StatusMovedPermanently)
	})
	mux.HandleFunc("GET /api/v1/seasons", sentryHandler.HandleFunc(seasonsHandler.ListSeasons))
	mux.Handle("GET /api/v1/seasons/current", sentryHandler.HandleFunc(seasonsHandler.Current))
	mux.HandleFunc("GET /api/v1/seasons/{id}/submissions", sentryHandler.HandleFunc(seasonSubmissionsHandler.ListSubmissions))
	mux.HandleFunc("POST /api/v1/seasons/{id}/submissions", sentryHandler.HandleFunc(seasonSubmissionsHandler.CreateSubmission))

	return middleware.Recover(middleware.IPFilter(middleware.Compression(mux)))
}
