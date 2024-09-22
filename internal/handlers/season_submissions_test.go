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

func TestCreateSeasonSubmissionHandler(t *testing.T) {

	t.Run("returns bad request error if body is not valid JSON", func(t *testing.T) {
		repo := repository.NewSeasonSubmissionsRepository(nil)
		submissionsHandler := handlers.NewSeasonSubmissionsHandler(repo)

		req, err := http.NewRequest("POST", "/api/v1/seasons/1/submissions", strings.NewReader(`not-json`))
		require.NoError(t, err)

		rr := httptest.NewRecorder()
		mux := http.NewServeMux()
		mux.HandleFunc("POST /api/v1/seasons/{id}/submissions", submissionsHandler.CreateSubmission)

		mux.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("returns an error if required fields are missing", func(t *testing.T) {
		repo := repository.NewSeasonSubmissionsRepository(nil)
		submissionsHandler := handlers.NewSeasonSubmissionsHandler(repo)

		req, err := http.NewRequest("POST", "/api/v1/seasons/1/submissions", strings.NewReader(`{}`))
		require.NoError(t, err)

		rr := httptest.NewRecorder()
		mux := http.NewServeMux()
		mux.HandleFunc("POST /api/v1/seasons/{id}/submissions", submissionsHandler.CreateSubmission)

		mux.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
		assert.Contains(t, rr.Body.String(), `"status":422`)
		assert.Contains(t, rr.Body.String(), `"param":"name"`)
		assert.Contains(t, rr.Body.String(), `"param":"class"`)
		assert.Contains(t, rr.Body.String(), `"param":"mode"`)
		assert.Contains(t, rr.Body.String(), `"param":"tier"`)
		assert.Contains(t, rr.Body.String(), `"param":"video"`)
		assert.Contains(t, rr.Body.String(), `"param":"duration"`)
	})
}
