package activitypub

import (
	"encoding/json"
	"fmt"
	rmodels "github.com/feditools/relay/internal/activitypub/models"
	"github.com/feditools/relay/internal/path"
	"github.com/tyrm/go-util/mimetype"
	"net/http"
)

func (m *Module) wellknownWebFingerGetHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "wellknownWebFingerGetHandler")

	subject := r.URL.Query().Get("resource")
	if subject != fmt.Sprintf("acct:relay@%s", m.domain) {
		w.Header().Set("Content-Type", mimetype.TextPlain)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("%d %s", http.StatusNotFound, http.StatusText(http.StatusNotFound))))
		return
	}

	webfinger := rmodels.WebFinger{
		Aliases: []string{path.GenActor(m.domain)},
		Links: []rmodels.Link{
			{
				Href: path.GenActor(m.domain),
				Rel:  "self",
				Type: mimetype.ApplicationActivityJSON,
			},
			{
				Href: path.GenActor(m.domain),
				Rel:  "self",
				Type: mimetype.ApplicationLDJSONActivityStreams,
			},
		},
		Subject: subject,
	}

	w.Header().Set("Content-Type", mimetype.ApplicationJSON)
	err := json.NewEncoder(w).Encode(webfinger)
	if err != nil {
		l.Errorf("marshaling json: %s", err.Error())
	}
}
