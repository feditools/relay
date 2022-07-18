package http

import (
	"github.com/go-http-utils/etag"
	"github.com/gorilla/handlers"
	"github.com/tyrm/go-util/middleware"
	"net/http"
	"time"
)

func (s *Server) MiddlewareMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		l := logger.WithField("func", "middlewareMetrics")

		wx := NewResponseWriter(w)

		// Do Request
		next.ServeHTTP(wx, r)

		ended := time.Since(start)
		l.Debugf("rendering %s took %d ms", r.URL.Path, ended.Milliseconds())
		go s.metrics.HTTPRequestTiming(ended, wx.Status(), r.Method, r.URL.Path)
	})
}

// WrapInMiddlewares wraps an http.Handler in the server's middleware.
func (s *Server) WrapInMiddlewares(h http.Handler) http.Handler {
	return s.MiddlewareMetrics(
		middleware.BlockMissingUserAgentMux(
			etag.Handler(
				handlers.CompressHandler(
					middleware.BlockFlocMux(
						h,
					),
				), false,
			),
		),
	)
}
