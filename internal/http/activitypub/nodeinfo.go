package activitypub

import (
	"encoding/json"
	"github.com/feditools/relay/internal/models"
	"github.com/feditools/relay/internal/path"
	"github.com/tyrm/go-util/mimetype"
	"net/http"
)

func (m *Module) nodeinfo20GetHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "nodeinfo20GetHandler")

	peers, err := m.logic.GetPeers(r.Context())
	if err != nil {
		l.Errorf("get peers: %s", err.Error())
		w.Header().Set("Content-Type", mimetype.TextPlain)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	nodeinfo := models.NodeInfo{
		Metadata: map[string]interface{}{
			"peers": peers,
		},
		OpenRegistrations: true,
		Protocols:         []string{"activitypub"},
		Services: models.Services{
			Inbound:  []string{},
			Outbound: []string{},
		},
		Software: models.Software{
			Name:    m.appName,
			Version: m.appVersion,
		},
		Usage: models.Usage{
			LocalPosts: 0,
			Users: models.UsageUsers{
				Total: 1,
			},
		},
		Version: "2.0",
	}

	w.Header().Set("Content-Type", mimetype.ApplicationJSON)
	err = json.NewEncoder(w).Encode(nodeinfo)
	if err != nil {
		l.Errorf("marshaling json: %s", err.Error())
	}
}

func (m *Module) wellknownNodeInfoGetHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "wellknownNodeInfoGetHandler")

	wellknown := models.NodeInfoWellKnown{
		Links: []models.Link{
			{
				Rel:  "http://nodeinfo.diaspora.software/ns/schema/2.0",
				Href: path.GenNodeinfo20(m.logic.Domain()),
			},
		},
	}

	w.Header().Set("Content-Type", mimetype.ApplicationJSON)
	err := json.NewEncoder(w).Encode(wellknown)
	if err != nil {
		l.Errorf("marshaling json: %s", err.Error())
	}
}
