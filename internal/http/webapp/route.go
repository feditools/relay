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

	webapp := s.PathPrefix("/").Subrouter()
	webapp.Use(m.Middleware)
	webapp.NotFoundHandler = m.notFoundHandler()
	webapp.MethodNotAllowedHandler = m.methodNotAllowedHandler()

	webapp.HandleFunc(path.Home, m.HomeGetHandler).Methods(nethttp.MethodGet)

	return nil
}
