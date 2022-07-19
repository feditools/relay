package webapp

import (
	"github.com/feditools/relay/internal/http"
	"github.com/feditools/relay/internal/path"
	"github.com/feditools/relay/web"
	"io/fs"
	nethttp "net/http"
)

// Route attaches routes to the web server.
func (m *Module) Route(s *http.Server) error {
	staticFS, err := fs.Sub(web.Files, DirStatic)
	if err != nil {
		return err
	}

	// Static Files
	s.PathPrefix(path.Static).Handler(nethttp.StripPrefix(path.Static, nethttp.FileServer(nethttp.FS(staticFS))))

	webapp := s.PathPrefix(path.App).Subrouter()
	webapp.Use(m.Middleware)
	webapp.NotFoundHandler = m.notFoundHandler()
	webapp.MethodNotAllowedHandler = m.methodNotAllowedHandler()

	webapp.HandleFunc(path.AppSubHome, m.HomeGetHandler).Methods(nethttp.MethodGet)
	webapp.HandleFunc(path.AppSubLogin, m.LoginGetHandler).Methods(nethttp.MethodGet)
	webapp.HandleFunc(path.AppSubLogin, m.LoginPostHandler).Methods(nethttp.MethodPost)

	admin := s.PathPrefix(path.AppAdmin).Subrouter()
	admin.NotFoundHandler = m.notFoundHandler()
	admin.MethodNotAllowedHandler = m.methodNotAllowedHandler()

	return nil
}
