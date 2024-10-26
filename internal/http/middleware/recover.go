package middleware

import (
	"fmt"
	"net/http"

	"github.com/janczizikow/pit/internal/http/response"
	"github.com/rs/zerolog/log"
)

// Recoverer is a middleware that recovers from panics, logs the error
// and returns a HTTP 500 (Internal Server Error) status.
func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				log.Error().Interface("error", err).Msg("recovered from panic")
				response.ErrorResponse(w, r, http.StatusInternalServerError, fmt.Sprintf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
