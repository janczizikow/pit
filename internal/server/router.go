package server

import (
	"fmt"
	"net/http"
	"os"

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

	path := os.Getenv("WEB_ROOT")
	fmt.Println(path)
	if path != "" {
		static := http.FileServer(http.Dir(path))
		mux.Handle("/", static)
		mux.Handle("/submission", static)
	}
	mux.HandleFunc("GET /api/v1/seasons", sentryHandler.HandleFunc(seasonsHandler.ListSeasons))
	mux.HandleFunc("GET /api/v1/seasons/{id}/submissions", seasonSubmissionsHandler.ListSubmissions)
	mux.HandleFunc("POST /api/v1/seasons/{id}/submissions", seasonSubmissionsHandler.CreateSubmission)

	return middleware.Recover(mux)
}
