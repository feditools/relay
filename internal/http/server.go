package http

import (
	"context"
	"github.com/feditools/relay/internal/config"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/tyrm/go-util/middleware"
	"net/http"
	"time"
)

// Server is a http 2 web server
type Server struct {
	router *mux.Router
	srv    *http.Server
}

// HandleFunc attaches a function to a path
func (r *Server) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	return r.router.HandleFunc(path, f)
}

// PathPrefix attaches a new route url path prefix
func (r *Server) PathPrefix(path string) *mux.Route {
	return r.router.PathPrefix(path)
}

// Start starts the web server
func (r *Server) Start() error {
	l := logger.WithField("func", "Start")
	l.Infof("listening on %s", r.srv.Addr)
	return r.srv.ListenAndServe()
}

// Stop shuts down the web server
func (r *Server) Stop(ctx context.Context) error {
	return r.srv.Shutdown(ctx)
}

// NewServer creates a new http web server
func NewServer(_ context.Context) (*Server, error) {
	r := mux.NewRouter()
	r.Use(handlers.CompressHandler)
	r.Use(middleware.BlockFlocMux)

	s := &http.Server{
		Addr:         viper.GetString(config.Keys.ServerHTTPBind),
		Handler:      r,
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}

	return &Server{
		router: r,
		srv:    s,
	}, nil
}
