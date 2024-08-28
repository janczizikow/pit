package response

import (
	"encoding/json"
	"net/http"
)

const contentType = "application/json; charset=utf-8"

// WriteJSON sends a JSON response with status code.
func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		writeErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(status)
	_, err = w.Write(bytes)
	if err != nil {
		writeErrorJSON(w, http.StatusInternalServerError, err)
	}
}

func writeErrorJSON(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(status)
	w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:golint,errcheck
}
