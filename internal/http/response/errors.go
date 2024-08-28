package response

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

type ValidationError struct {
	Status int                 `json:"status"`
	Errors []map[string]string `json:"errors"`
}

func ErrorResponse(w http.ResponseWriter, r *http.Request, status int, message string) {
	log.Error().Str("request_method", r.Method).Str("request_url", r.URL.String()).Msg(message)
	data := map[string]interface{}{"status": status, "error": message}
	WriteJSON(w, status, data)
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	ErrorResponse(w, r, http.StatusBadRequest, err.Error())
}

func FailedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	log.Error().Str("request_method", r.Method).Str("request_url", r.URL.String()).Msg("failed validation")
	data := ValidationError{Status: http.StatusUnprocessableEntity}
	for k, v := range errors {
		data.Errors = append(data.Errors, map[string]string{"param": k, "error": v})
	}
	WriteJSON(w, http.StatusUnprocessableEntity, data)
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	ErrorResponse(w, r, http.StatusNotFound, "Not found")
}

func MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	ErrorResponse(w, r, http.StatusMethodNotAllowed, "Method not allowed")
}

func InternalServerErrorResponse(w http.ResponseWriter, r *http.Request) {
	ErrorResponse(w, r, http.StatusInternalServerError, "Internal Server Error")
}
