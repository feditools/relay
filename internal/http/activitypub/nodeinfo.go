package activitypub

import (
	"encoding/json"
	"github.com/feditools/go-lib/fedihelper"
	libhttp "github.com/feditools/go-lib/http"
	"github.com/feditools/relay/internal/path"
	"net/http"
)

func (m *Module) nodeinfo20GetHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "nodeinfo20GetHandler")

	peers, err := m.logic.GetPeers(r.Context())
	if err != nil {
		l.Errorf("get peers: %s", err.Error())
		w.Header().Set("Content-Type", libhttp.MimeTextPlain.String())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	nodeinfo := fedihelper.NodeInfoV2{
		Metadata: map[string]interface{}{
			"peers": peers,
		},
		OpenRegistrations: true,
		Protocols:         []string{"activitypub"},
		Services: fedihelper.Services{
			Inbound:  []string{},
			Outbound: []string{},
		},
		Software: fedihelper.Software{
			Name:    m.appName,
			Version: m.appVersion,
		},
		Usage: fedihelper.Usage{
			LocalPosts: 0,
			Users: fedihelper.UsageUsers{
				Total: 1,
			},
		},
		Version: "2.0",
	}

	w.Header().Set("Content-Type", libhttp.MimeAppJSON.String())
	err = json.NewEncoder(w).Encode(nodeinfo)
	if err != nil {
		l.Errorf("marshaling json: %s", err.Error())
	}
}

func (m *Module) wellknownNodeInfoGetHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "wellknownNodeInfoGetHandler")

	wellknown := fedihelper.WellKnownNodeInfo{
		Links: []fedihelper.Link{
			{
				Rel:  "http://nodeinfo.diaspora.software/ns/schema/2.0",
				Href: path.GenNodeinfo20(m.logic.Domain()),
			},
		},
	}

	w.Header().Set("Content-Type", libhttp.MimeAppJSON.String())
	err := json.NewEncoder(w).Encode(wellknown)
	if err != nil {
		l.Errorf("marshaling json: %s", err.Error())
	}
}
