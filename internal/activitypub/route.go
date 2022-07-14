package activitypub

import (
	"github.com/feditools/relay/internal/http"
	"github.com/feditools/relay/internal/path"
)

// Route attaches routes to the web server
func (m *Module) Route(s *http.Server) error {
	ap := s.PathPrefix("/").Subrouter()
	ap.Use(m.middlewareCheckHTTPSig)
	ap.HandleFunc(path.APActor, m.actorGetHandler).Methods("GET")
	ap.HandleFunc(path.APInbox, m.inboxPostHandler).Methods("POST")
	ap.HandleFunc(path.APNodeInfo20, m.nodeinfo20GetHandler).Methods("GET")
	ap.HandleFunc(path.APWellKnownNodeInfo, m.wellknownNodeInfoGetHandler).Methods("GET")
	ap.HandleFunc(path.APWellKnownWebFinger, m.wellknownWebFingerGetHandler).Methods("GET")
	return nil
}
