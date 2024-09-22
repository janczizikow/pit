package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/janczizikow/pit/internal/handlers"
	"github.com/janczizikow/pit/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListSeasonsHandler(t *testing.T) {
	t.Run("returns 200 for a valid request", func(t *testing.T) {
		repo := repository.NewSeasonsRepository(db)
		seasonsHandler := handlers.NewSeasonsHandler(repo)

		req, err := http.NewRequest("GET", "/api/v1/seasons", strings.NewReader(``))
		require.NoError(t, err)

		rr := httptest.NewRecorder()
		mux := http.NewServeMux()
		mux.HandleFunc("GET /api/v1/seasons", seasonsHandler.ListSeasons)

		mux.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})
}
