package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/janczizikow/pit/internal/handlers"
	"github.com/janczizikow/pit/internal/repository"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateSubmissionHandler(t *testing.T) {
	t.Run("returns bad request error if body is not valid JSON", func(t *testing.T) {
		repo := repository.NewSubmissionsRepository(nil)
		submissionsHandler := handlers.NewSubmissionsHandler(repo)

		req, err := http.NewRequest("POST", "/api/v1/submissions", strings.NewReader(`not-json`))
		require.NoError(t, err)

		rr := httptest.NewRecorder()
		router := httprouter.New()
		router.POST("/api/v1/submissions", submissionsHandler.CreateSubmission)

		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("returns an error if required fields are missing", func(t *testing.T) {
		repo := repository.NewSubmissionsRepository(nil)
		submissionsHandler := handlers.NewSubmissionsHandler(repo)

		req, err := http.NewRequest("POST", "/api/v1/submissions", strings.NewReader(`{}`))
		require.NoError(t, err)

		rr := httptest.NewRecorder()
		router := httprouter.New()
		router.POST("/api/v1/submissions", submissionsHandler.CreateSubmission)

		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
		assert.Contains(t, rr.Body.String(), `"status":422`)
		assert.Contains(t, rr.Body.String(), `"param":"name"`)
		assert.Contains(t, rr.Body.String(), `"param":"class"`)
		assert.Contains(t, rr.Body.String(), `"param":"mode"`)
		assert.Contains(t, rr.Body.String(), `"param":"tier"`)
		assert.Contains(t, rr.Body.String(), `"param":"video"`)
		assert.Contains(t, rr.Body.String(), `"param":"duration"`)
	})

	t.Run("returns 201 if data is valid", func(t *testing.T) {
		repo := repository.NewSubmissionsRepository(db)
		submissionsHandler := handlers.NewSubmissionsHandler(repo)

		req, err := http.NewRequest("POST", "/api/v1/submissions", strings.NewReader(`{
			"name": "Test",
			"class": "rogue",
			"mode": "softcore",
			"tier": 150,
			"duration": 300,
			"video": "https://youtube.com"
		}`))
		require.NoError(t, err)

		rr := httptest.NewRecorder()
		router := httprouter.New()
		router.POST("/api/v1/submissions", submissionsHandler.CreateSubmission)

		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusCreated, rr.Code)
	})
}
