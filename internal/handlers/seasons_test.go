package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/janczizikow/pit/internal/handlers"
	"github.com/janczizikow/pit/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCurrentSeasonHandler(t *testing.T) {
	t.Run("returns 404 when no seasons exist", func(t *testing.T) {
		repo := repository.NewSeasonsRepository(db)
		seasonsHandler := handlers.NewSeasonsHandler(repo)

		req, err := http.NewRequest(http.MethodGet, "/api/v1/seasons/current", nil)
		require.NoError(t, err)

		rr := httptest.NewRecorder()
		mux := http.NewServeMux()
		mux.HandleFunc("GET /api/v1/seasons/current", seasonsHandler.Current)

		mux.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.JSONEq(t, rr.Body.String(), `{"error":"Not found", "status":404}`)
	})
}

func TestListSeasonsHandler(t *testing.T) {
	t.Run("returns 200 for a valid request", func(t *testing.T) {
		repo := repository.NewSeasonsRepository(db)
		seasonsHandler := handlers.NewSeasonsHandler(repo)

		req, err := http.NewRequest("GET", "/api/v1/seasons", nil)
		require.NoError(t, err)

		rr := httptest.NewRecorder()
		mux := http.NewServeMux()
		mux.HandleFunc("GET /api/v1/seasons", seasonsHandler.ListSeasons)

		mux.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})
}
