package server

import (
	"net/http"
	"os"

	"github.com/janczizikow/pit/internal/handlers"
	"github.com/janczizikow/pit/internal/http/middleware"
	"github.com/janczizikow/pit/internal/repository"
)

func Router(s *Server) http.Handler {
	mux := http.NewServeMux()

	repo := repository.New(s.db)
	submissionsHandler := handlers.NewSubmissionsHandler(repo.Submissions)
	path := os.Getenv("WEB_ROOT")
	static := http.FileServer(http.Dir(path))
	mux.Handle("/", static)
	mux.Handle("/submission", static)
	mux.HandleFunc("GET /api/v1/submissions", submissionsHandler.ListSubmissions)
	mux.HandleFunc("POST /api/v1/submissions", submissionsHandler.CreateSubmission)

	return middleware.Recover(mux)
}
