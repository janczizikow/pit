package handlers

import (
	"net/http"

	"github.com/janczizikow/pit/internal/http/request"
	"github.com/janczizikow/pit/internal/http/response"
	"github.com/janczizikow/pit/internal/models"
	"github.com/janczizikow/pit/internal/repository"
	"github.com/janczizikow/pit/internal/validator"
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
	page, err := request.QueryInt(r, "page", 1)
	if err != nil {
		response.BadRequestResponse(w, r, err)
		return
	}
	size, err := request.QueryInt(r, "size", 100)
	if err != nil {
		response.BadRequestResponse(w, r, err)
		return
	}
	sort, err := request.QueryStrings(r, "sort", "")
	if err != nil {
		response.BadRequestResponse(w, r, err)
		return
	}
	class, err := request.QueryString(r, "class", "")
	if err != nil {
		response.BadRequestResponse(w, r, err)
		return
	}

	paginator := request.NewPaginator(size, page, sort, map[string]bool{
		"duration":  true,
		"tier":      true,
		"-duration": true,
		"-tier":     true,
	})
	if ok, errs := paginator.Valid(); !ok {
		response.FailedValidationResponse(w, r, errs)
		return
	}

	submissions, total, err := h.repo.List(paginator.Limit(), paginator.Offset(), class, paginator.Sort())
	if err != nil {
		// TODO: handle enum errors
		response.InternalServerErrorResponse(w, r)
		return
	}

	response.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"data":     submissions,
		"metadata": paginator.CalculateMetadata(total),
	})
}

func (h *submissionsHandler) CreateSubmission(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	submission := &models.Submission{}
	err := request.ReadJSON(w, r, submission)
	if err != nil {
		response.BadRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	if models.ValidateSubmission(v, submission); !v.Valid() {
		response.FailedValidationResponse(w, r, v.Errors)
		return
	}

	created, err := h.repo.Create(submission)
	if err != nil {
		response.InternalServerErrorResponse(w, r)
		return
	}
	response.WriteJSON(w, http.StatusCreated, created)
}
