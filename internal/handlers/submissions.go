package handlers

import (
	"net/http"

	"github.com/janczizikow/pit/internal/repository"
	"github.com/julienschmidt/httprouter"
)

type SubmissionsHandler interface {
	ListSubmissions(w http.ResponseWriter, r *http.Request)
	CreateSubmission(w http.ResponseWriter, r *http.Request)
}

type submissionsHandler struct {
	repo repository.SubmissionsRepository
}

func NewSubmissionsHandler(repo repository.SubmissionsRepository) *submissionsHandler {
	return &submissionsHandler{repo: repo}
}

func (h *submissionsHandler) ListSubmissions(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
}

func (h *submissionsHandler) CreateSubmission(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
}
