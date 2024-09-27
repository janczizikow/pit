package handlers_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/janczizikow/pit/internal/handlers"
	"github.com/janczizikow/pit/internal/models"
	"github.com/janczizikow/pit/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/rand"
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

	t.Run("returns 201 if data is valid", func(t *testing.T) {
		repo := repository.NewSeasonSubmissionsRepository(db)
		submissionsHandler := handlers.NewSeasonSubmissionsHandler(repo)
		seasonId, err := createSeason(db)
		require.NoError(t, err)

		req, err := http.NewRequest("POST", fmt.Sprintf("/api/v1/seasons/%v/submissions", seasonId), strings.NewReader(`{
			"name": "Test",
			"class": "rogue",
			"mode": "softcore",
			"tier": 150,
			"duration": 300,
			"video": "https://youtube.com"
		}`))
		require.NoError(t, err)

		rr := httptest.NewRecorder()
		mux := http.NewServeMux()
		mux.HandleFunc("POST /api/v1/seasons/{id}/submissions", submissionsHandler.CreateSubmission)

		t.Cleanup(func() {
			submission := models.Submission{}
			err := json.Unmarshal(rr.Body.Bytes(), &submission)
			require.NoError(t, err)
			_, err = db.Exec(context.Background(), "DELETE FROM submissions WHERE id = $1", submission.ID)
			require.NoError(t, err)
			_, err = db.Exec(context.Background(), "DELETE FROM seasons WHERE id = $1", seasonId)
			require.NoError(t, err)
		})
		mux.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusCreated, rr.Code)
	})
}

func TestListSeasonSubmissionHandler(t *testing.T) {
	repo := repository.NewSeasonSubmissionsRepository(db)
	submissionsHandler := handlers.NewSeasonSubmissionsHandler(repo)
	seasonId, err := createSeason(db)
	require.NoError(t, err)

	t.Cleanup(func() {
		_, err = db.Exec(context.Background(), "DELETE FROM seasons WHERE id = $1", seasonId)
		require.NoError(t, err)
	})

	t.Run("returns 200 if no submissions exist", func(t *testing.T) {
		req, err := http.NewRequest("GET", fmt.Sprintf("/api/v1/seasons/%v/submissions", seasonId), nil)
		require.NoError(t, err)

		rr := httptest.NewRecorder()
		mux := http.NewServeMux()
		mux.HandleFunc("GET /api/v1/seasons/{id}/submissions", submissionsHandler.ListSubmissions)

		mux.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.JSONEq(t, rr.Body.String(), `{"data":[],"metadata":{}}`)
	})

	t.Run("returns 200 and empty array if class query is invalid", func(t *testing.T) {
		req, err := http.NewRequest("GET", fmt.Sprintf("/api/v1/seasons/%v/submissions?class=test", seasonId), nil)
		require.NoError(t, err)

		rr := httptest.NewRecorder()
		mux := http.NewServeMux()
		mux.HandleFunc("GET /api/v1/seasons/{id}/submissions", submissionsHandler.ListSubmissions)

		mux.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.JSONEq(t, rr.Body.String(), `{"data":[],"metadata":{}}`)
	})

	t.Run("returns 422 if pagination is invalid", func(t *testing.T) {
		req, err := http.NewRequest("GET", fmt.Sprintf("/api/v1/seasons/%v/submissions?page=-1&size=1001", seasonId), nil)
		require.NoError(t, err)

		rr := httptest.NewRecorder()
		mux := http.NewServeMux()
		mux.HandleFunc("GET /api/v1/seasons/{id}/submissions", submissionsHandler.ListSubmissions)

		mux.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	})
}

func createSeason(db *pgxpool.Pool) (int, error) {
	var id int
	err := db.QueryRow(
		context.Background(),
		`INSERT INTO seasons (name, pit, start)
		 VALUES ($1, $2, $3)
		 RETURNING id;`,
		randStr(36),
		false,
		time.Now(),
	).Scan(&id)
	return id, err
}

func randStr(length uint) string {
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(uint64(time.Now().UnixNano())))
	bytes := make([]byte, int(length))
	for i := uint(0); i < length; i++ {
		bytes[i] = byte('a' + seededRand.Intn('z'-'a'))
	}
	return string(bytes)
}
