package handlers

import (
	"net/http"

	"github.com/janczizikow/pit/internal/http/request"
	"github.com/janczizikow/pit/internal/http/response"
	"github.com/janczizikow/pit/internal/repository"
)

type SeasonsHandler interface {
	ListSeasons(w http.ResponseWriter, r *http.Request)
	Current(w http.ResponseWriter, r *http.Request)
}

type seasonsHandler struct {
	repo repository.SeasonsRepository
}

func NewSeasonsHandler(repo repository.SeasonsRepository) SeasonsHandler {
	return &seasonsHandler{repo: repo}
}

func (h *seasonsHandler) ListSeasons(w http.ResponseWriter, r *http.Request) {
	paginator := request.NewPaginator(100, 1, make([]string, 0), make(map[string]bool))
	seasons, total, err := h.repo.List()
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
