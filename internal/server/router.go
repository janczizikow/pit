package server

import (
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
	submissionsHandler := handlers.NewSubmissionsHandler(repo.Submissions)

	path := os.Getenv("WEB_ROOT")

	static := http.FileServer(http.Dir(path))
	mux.Handle("/", static)
	mux.Handle("/submission", static)
	mux.HandleFunc("GET /api/v1/submissions", sentryHandler.HandleFunc(submissionsHandler.ListSubmissions))
	mux.HandleFunc("POST /api/v1/submissions", sentryHandler.HandleFunc(submissionsHandler.CreateSubmission))

	return middleware.Recover(mux)
}
