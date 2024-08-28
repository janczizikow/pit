package response_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/janczizikow/pit/internal/http/response"
	"github.com/stretchr/testify/assert"
)

func TestWriteJSON(t *testing.T) {
	testData := map[string]interface{}{
		"key": "value",
	}
	rr := httptest.NewRecorder()
	response.WriteJSON(rr, http.StatusOK, testData)

	assert.Equal(t, "application/json; charset=utf-8", rr.Header().Get("Content-Type"))
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, `{"key":"value"}`, rr.Body.String())
}

func TestWriteJSONError(t *testing.T) {
	w := &mockResponseWriter{writeError: true}
	response.WriteJSON(w, http.StatusOK, nil)

	assert.Equal(t, http.StatusInternalServerError, w.statusCode)
	assert.Equal(t, `{"error": "mocked error"}`, w.body)
}

// mockResponseWriter is a mock implementation of http.ResponseWriter for testing purposes
type mockResponseWriter struct {
	statusCode int
	body       string
	writeError bool
}

func (m *mockResponseWriter) Header() http.Header {
	return make(http.Header)
}

func (m *mockResponseWriter) WriteHeader(statusCode int) {
	m.statusCode = statusCode
}

func (m *mockResponseWriter) Write(body []byte) (int, error) {
	m.body = string(body)
	if m.writeError {
		return 0, errors.New("mocked error")
	}
	return len(body), nil
}
