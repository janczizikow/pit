package request

import (
	"net/http"
	"strconv"
)

// QueryParameter extracts a given key from query parameters.
// Returns an empty string if parameters wasn't found
func QueryParameter(r *http.Request, key string) string {
	if r == nil || r.URL == nil {
		return ""
	}
	return r.URL.Query().Get(key)
}

// QueryInt parses an integer from a query parameter.
// Returns defaultValue if the query parameter was empty, because query parameters are always optional.
func QueryInt(r *http.Request, key string, defaultValue int) (int, error) {
	s := QueryParameter(r, key)
	if s == "" {
		return defaultValue, nil
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return defaultValue, err
	}

	return i, nil
}
