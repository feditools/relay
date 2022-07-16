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
	if subject != fmt.Sprintf("acct:relay@%s", m.logic.Domain()) {
		w.Header().Set("Content-Type", mimetype.TextPlain)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	webfinger := rmodels.WebFinger{
		Aliases: []string{path.GenActor(m.logic.Domain())},
		Links: []rmodels.Link{
			{
				Href: path.GenActor(m.logic.Domain()),
				Rel:  "self",
				Type: mimetype.ApplicationActivityJSON,
			},
			{
				Href: path.GenActor(m.logic.Domain()),
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
