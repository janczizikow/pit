package middleware

import "net/http"

// CORS is a middleware that sets
// the Access-Control-Allow-Origin and Access-Control-Allow-Headers headers
// to "*" for all requests.
// This allows all origins to access the API.
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		next.ServeHTTP(w, r)
	})
}
