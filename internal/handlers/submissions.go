package handlers

import (
	"net/http"

	"github.com/janczizikow/pit/internal/http/request"
	"github.com/janczizikow/pit/internal/http/response"
	"github.com/janczizikow/pit/internal/models"
	"github.com/janczizikow/pit/internal/repository"
	"github.com/janczizikow/pit/internal/validator"
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

func (h *submissionsHandler) ListSubmissions(w http.ResponseWriter, r *http.Request) {
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
	mode, err := request.QueryString(r, "mode", "")
	if err != nil {
		response.BadRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	v.Check(validator.In(class, "", models.Barbarian, models.Druid, models.Necromancer, models.Rogue, models.Sorcerer), "class", "is invalid")
	v.Check(validator.In(mode, "", models.Softcore, models.Hardcore), "mode", "is invalid")

	paginator := request.NewPaginator(size, page, sort, map[string]bool{
		"created_at":  true,
		"duration":    true,
		"tier":        true,
		"-created_at": true,
		"-duration":   true,
		"-tier":       true,
	})
	if ok, errs := paginator.Valid(); !ok {
		response.FailedValidationResponse(w, r, errs)
		return
	}

	submissions, total, err := h.repo.List(
		repository.ListSubmissionsParams{
			Class:   class,
			Mode:    mode,
			OrderBy: paginator.Sort(),
			Limit:   paginator.Limit(),
			Offset:  paginator.Offset(),
		},
	)
	if err != nil {
		response.InternalServerErrorResponse(w, r)
		return
	}

	response.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"data":     submissions,
		"metadata": paginator.CalculateMetadata(total),
	})
}

func (h *submissionsHandler) CreateSubmission(w http.ResponseWriter, r *http.Request) {
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
