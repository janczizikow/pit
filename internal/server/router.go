package server

import (
	"net/http"

	"github.com/janczizikow/pit/internal/handlers"
	"github.com/janczizikow/pit/internal/http/middleware"
	"github.com/janczizikow/pit/internal/repository"
)

func Router(s *Server) http.Handler {
	mux := http.NewServeMux()

	repo := repository.New(s.db)
	submissionsHandler := handlers.NewSubmissionsHandler(repo.Submissions)

	mux.Handle("/", http.FileServer(http.Dir("./web/build")))
	mux.Handle("/submission", http.FileServer(http.Dir("./web/build")))
	mux.HandleFunc("GET /api/v1/submissions", submissionsHandler.ListSubmissions)
	mux.HandleFunc("POST /api/v1/submissions", submissionsHandler.CreateSubmission)

	return middleware.Recover(mux)
}
