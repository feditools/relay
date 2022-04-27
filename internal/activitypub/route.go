package activitypub

import (
	"github.com/feditools/relay/internal/http"
	"github.com/feditools/relay/internal/path"
)

// Route attaches routes to the web server
func (m *Module) Route(s *http.Server) error {
	s.HandleFunc(path.APActor, m.actorGetHandler).Methods("GET")
	return nil
}
