package webapp

import (
	"github.com/feditools/relay/internal/http"
	"github.com/feditools/relay/internal/path"
	nethttp "net/http"
)

// Route attaches routes to the web server.
func (m *Module) Route(s *http.Server) error {
	s.HandleFunc("/", m.ForwardToHomeHandler).Methods(nethttp.MethodGet)
	s.HandleFunc(path.App, m.ForwardToHomeHandler).Methods(nethttp.MethodGet)

	webapp := s.PathPrefix(path.App).Subrouter()
	webapp.Use(m.Middleware)
	webapp.NotFoundHandler = m.notFoundHandler()
	webapp.MethodNotAllowedHandler = m.methodNotAllowedHandler()

	webapp.HandleFunc(path.AppSubHome, m.HomeGetHandler).Methods(nethttp.MethodGet)
	webapp.HandleFunc(path.AppSubLogin, m.LoginGetHandler).Methods(nethttp.MethodGet)
	webapp.HandleFunc(path.AppSubLogin, m.LoginPostHandler).Methods(nethttp.MethodPost)
	webapp.HandleFunc(path.AppSubLogout, m.LogoutGetHandler).Methods(nethttp.MethodGet)

	admin := s.PathPrefix(path.AppAdmin).Subrouter()
	admin.NotFoundHandler = m.notFoundHandler()
	admin.MethodNotAllowedHandler = m.methodNotAllowedHandler()

	return nil
}
