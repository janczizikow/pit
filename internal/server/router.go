package server

import (
	"net/http"
	"os"

	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/janczizikow/pit/internal/handlers"
	"github.com/janczizikow/pit/internal/http/middleware"
	"github.com/janczizikow/pit/internal/repository"
	"github.com/klauspost/compress/gzhttp"
)

func Router(s *Server) http.Handler {
	mux := http.NewServeMux()
	sentryHandler := sentryhttp.New(sentryhttp.Options{
		Repanic: true,
	})

	repo := repository.New(s.db)
	seasonsHandler := handlers.NewSeasonsHandler(repo.Seasons)
	seasonSubmissionsHandler := handlers.NewSeasonSubmissionsHandler(repo.SeasonSubmissions)

	path := os.Getenv("WEB_ROOT")
	if path != "" {
		static := http.FileServer(http.Dir(path))
		mux.Handle("/", gzhttp.GzipHandler(static))
		mux.Handle("/submission", gzhttp.GzipHandler(static))
	}
	mux.HandleFunc("GET /api/v1/seasons", sentryHandler.HandleFunc(seasonsHandler.ListSeasons))
	mux.HandleFunc("GET /api/v1/seasons/{id}/submissions", seasonSubmissionsHandler.ListSubmissions)
	mux.HandleFunc("POST /api/v1/seasons/{id}/submissions", seasonSubmissionsHandler.CreateSubmission)

	return middleware.Recover(mux)
}
