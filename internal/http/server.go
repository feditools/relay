package http

import (
	"context"
	"fmt"
	libhttp "github.com/feditools/go-lib/http"
	"github.com/feditools/go-lib/metrics"
	"github.com/feditools/relay/internal/config"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

// Server is a http 2 web server
type Server struct {
	metrics metrics.Collector
	router  *mux.Router
	srv     *http.Server
}

// NewServer creates a new http web server
func NewServer(_ context.Context, m metrics.Collector) (*Server, error) {
	r := mux.NewRouter()

	s := &http.Server{
		Addr:         viper.GetString(config.Keys.ServerHTTPBind),
		Handler:      r,
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}

	server := &Server{
		metrics: m,
		router:  r,
		srv:     s,
	}

	// add global middlewares
	r.Use(server.WrapInMiddlewares)

	r.NotFoundHandler = server.notFoundHandler()
	r.MethodNotAllowedHandler = server.methodNotAllowedHandler()

	return server, nil
}

// HandleFunc attaches a function to a path
func (s *Server) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	return s.router.HandleFunc(path, f)
}

// PathPrefix attaches a new route url path prefix
func (s *Server) PathPrefix(path string) *mux.Route {
	return s.router.PathPrefix(path)
}

// Start starts the web server
func (s *Server) Start() error {
	l := logger.WithField("func", "Start")
	l.Infof("listening on %s", s.srv.Addr)
	return s.srv.ListenAndServe()
}

// Stop shuts down the web server
func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func (s *Server) methodNotAllowedHandler() http.Handler {
	// wrap in middleware since middlware isn't run on error pages
	return s.WrapInMiddlewares(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", libhttp.MimeTextPlain.String())
		w.Write([]byte(fmt.Sprintf("%d %s", http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))))
	}))
}

func (s *Server) notFoundHandler() http.Handler {
	// wrap in middleware since middlware isn't run on error pages
	return s.WrapInMiddlewares(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", libhttp.MimeTextPlain.String())
		w.Write([]byte(fmt.Sprintf("%d %s", http.StatusNotFound, http.StatusText(http.StatusNotFound))))
	}))
}
