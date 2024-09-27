package middleware_test

import (
	"compress/gzip"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/janczizikow/pit/internal/http/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompression(t *testing.T) {
	t.Parallel()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test response"))
	})

	req, err := http.NewRequest("GET", "/", nil)
	require.NoError(t, err)

	// Test without Accept-Encoding header
	rr := httptest.NewRecorder()
	compressionMiddleware := middleware.Compression(testHandler)
	compressionMiddleware.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "", rr.Header().Get("Content-Encoding"))
	assert.Equal(t, "test response", rr.Body.String())

	// Test with Accept-Encoding: gzip header
	req.Header.Set("Accept-Encoding", "gzip")
	rr = httptest.NewRecorder()
	compressionMiddleware.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "gzip", rr.Header().Get("Content-Encoding"))

	// Decompress the response body
	gzr, err := gzip.NewReader(rr.Body)
	require.NoError(t, err)
	defer gzr.Close()

	body, err := io.ReadAll(gzr)
	require.NoError(t, err)

	assert.Equal(t, "test response", string(body))
}
