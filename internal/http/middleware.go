package http

import (
	libhttp "github.com/feditools/go-lib/http"
	"github.com/go-http-utils/etag"
	"github.com/gorilla/handlers"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (s *Server) MiddlewareMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metric := s.metrics.NewHTTPRequest(r.Method, r.URL.Path)
		l := logger.WithFields(logrus.Fields{
			"func":      "middlewareMetrics",
			"client":    r.RemoteAddr,
			"useragent": r.UserAgent(),
		})

		wx := NewResponseWriter(w)

		// Do Request
		next.ServeHTTP(wx, r)

		go func() {
			ended := metric.Done(wx.Status())
			l.Debugf("rendering %s took %d ms", r.URL.Path, ended.Milliseconds())
		}()
	})
}

// WrapInMiddlewares wraps an http.Handler in the server's middleware.
func (s *Server) WrapInMiddlewares(h http.Handler) http.Handler {
	return s.MiddlewareMetrics(
		libhttp.BlockMissingUserAgent(
			libhttp.BlockFloc(
				etag.Handler(
					handlers.CompressHandler(
						h,
					), false,
				),
			),
		),
	)
}
