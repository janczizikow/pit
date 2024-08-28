package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/janczizikow/pit/internal/http/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRecover(t *testing.T) {
	panicHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("test panic")
	})

	req, err := http.NewRequest("GET", "/", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	recoverMiddleware := middleware.Recover(panicHandler)

	recoverMiddleware.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Contains(t, rr.Body.String(), "test panic")

	// Check if the log message contains the expected panic message
	// logMsg := log.Last().Message
	// expectedLogMsg := "recovered from panic"
	// if logMsg != expectedLogMsg {
	// 	t.Errorf("Expected log message '%s', got '%s'", expectedLogMsg, logMsg)
	// }
}
