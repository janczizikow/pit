package handlers

import (
	"net/http"
	"strconv"

	"github.com/janczizikow/pit/internal/http/request"
	"github.com/janczizikow/pit/internal/http/response"
	"github.com/janczizikow/pit/internal/repository"
)

type SeasonsHandler interface {
	ListSeasons(w http.ResponseWriter, r *http.Request)
	Current(w http.ResponseWriter, r *http.Request)
	GetStatistics(w http.ResponseWriter, r *http.Request)
}

type seasonsHandler struct {
	repo repository.SeasonsRepository
}

func NewSeasonsHandler(repo repository.SeasonsRepository) SeasonsHandler {
	return &seasonsHandler{repo: repo}
}

func (h *seasonsHandler) ListSeasons(w http.ResponseWriter, r *http.Request) {
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
	paginator := request.NewPaginator(size, page, make([]string, 0), make(map[string]bool))
	if ok, errs := paginator.Valid(); !ok {
		response.FailedValidationResponse(w, r, errs)
		return
	}
	seasons, total, err := h.repo.List(paginator.Limit(), paginator.Offset())
	if err != nil {
		response.InternalServerErrorResponse(w, r)
		return
	}
	response.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"data":     seasons,
		"metadata": paginator.CalculateMetadata(total),
	})
}

func (h *seasonsHandler) Current(w http.ResponseWriter, r *http.Request) {
	season, err := h.repo.Current()
	if err != nil {
		response.NotFoundResponse(w, r)
		return
	}
	response.WriteJSON(w, http.StatusOK, season)
}

func (h *seasonsHandler) GetStatistics(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	seasonId, err := strconv.Atoi(id)
	if err != nil || seasonId < 0 {
		response.NotFoundResponse(w, r)
		return
	}
	totals, statistics, err := h.repo.Statistics(seasonId)
	if err != nil {
		response.InternalServerErrorResponse(w, r)
		return
	}

	response.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"totals": totals,
		"data":   statistics,
	})
}
