package response

import (
	"encoding/json"
	"net/http"
)

const contentType = "application/json; charset=utf-8"

// WriteJSON sends a JSON response with status code.
func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", contentType)
	bytes, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeErrorJSON(w, err)
		return
	}
	w.WriteHeader(status)
	_, err = w.Write(bytes)
	if err != nil {
		writeErrorJSON(w, err)
	}
}

func writeErrorJSON(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:golint,errcheck
}
