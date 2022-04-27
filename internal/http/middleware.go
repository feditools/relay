package http

import (
	"net/http"
	"time"
)

func (s *Server) middlewareMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := logger.WithField("func", "middlewareMetrics")
		start := time.Now()

		wx := NewResponseWriter(w)

		// Do Request
		next.ServeHTTP(wx, r)

		ended := time.Since(start)
		l.Debugf("rendering %s took %d ms", r.URL.Path, ended.Milliseconds())
		go s.metrics.HTTPRequestTiming(ended, wx.Status(), r.Method, r.URL.Path)
	})
}
