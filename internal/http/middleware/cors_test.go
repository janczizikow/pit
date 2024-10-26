package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/janczizikow/pit/internal/http/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCORS(t *testing.T) {
	t.Parallel()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req, err := http.NewRequest("GET", "/", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	corsMiddleware := middleware.CORS(testHandler)

	corsMiddleware.ServeHTTP(rr, req)

	assert.Equal(t, "*", rr.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "*", rr.Header().Get("Access-Control-Allow-Headers"))
	assert.Equal(t, http.StatusOK, rr.Code)
}
