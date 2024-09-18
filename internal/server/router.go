package server

import (
	"net/http"

	"github.com/janczizikow/pit/internal/handlers"
	"github.com/janczizikow/pit/internal/http/middleware"
	"github.com/janczizikow/pit/internal/http/response"
	"github.com/janczizikow/pit/internal/repository"
	"github.com/julienschmidt/httprouter"
)

func Router(s *Server) http.Handler {
	router := httprouter.New()

	router.NotFound = http.FileServer(http.Dir("./web/build/"))
	router.MethodNotAllowed = http.HandlerFunc(response.MethodNotAllowedResponse)

	repo := repository.New(s.db)
	submissionsHandler := handlers.NewSubmissionsHandler(repo.Submissions)

	router.GET("/api/v1/submissions", submissionsHandler.ListSubmissions)
	router.POST("/api/v1/submissions", submissionsHandler.CreateSubmission)

	return middleware.Recover(router)
}
