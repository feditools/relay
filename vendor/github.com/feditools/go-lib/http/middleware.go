package http

import "net/http"

// BlockFloc is a middleware that applies a header to tell Google Chrome to disable cohort tracking.
func BlockFloc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Permissions-Policy", "interest-cohort=()")
		next.ServeHTTP(w, r)
	})
}

// BlockMissingUserAgent is a middleware that returns 400 Bad Request if the useragent is missing
func BlockMissingUserAgent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ua := r.UserAgent(); ua == "" {
			logger.WithField("func", "BlockMissingUserAgent").Debug("blocked request with missing user agent.")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
