package middleware

import "net/http"

// BlockMissingUserAgentMux is a mux middleware that returns 400 Bad Request if the useragent is missing
func BlockMissingUserAgentMux(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ua := r.UserAgent(); ua == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
