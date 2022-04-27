package activitypub

import (
	"github.com/feditools/relay/internal/http"
	"github.com/feditools/relay/internal/path"
)

// Route attaches routes to the web server
func (m *Module) Route(s *http.Server) error {
	s.HandleFunc(path.APActor, m.actorGetHandler).Methods("GET")
	s.HandleFunc(path.APNodeInfo20, m.nodeinfo20GetHandler).Methods("GET")
	s.HandleFunc(path.APWellKnownNodeInfo, m.wellknownNodeinfoGetHandler).Methods("GET")
	return nil
}
