package activitypub

import (
	"encoding/json"
	"fmt"
	"github.com/feditools/go-lib/fedihelper"
	libhttp "github.com/feditools/go-lib/http"
	"github.com/feditools/relay/internal/path"
	"net/http"
)

func (m *Module) wellknownWebFingerGetHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "wellknownWebFingerGetHandler")

	subject := r.URL.Query().Get("resource")
	if subject != fmt.Sprintf("acct:relay@%s", m.logic.Domain()) {
		w.Header().Set("Content-Type", libhttp.MimeTextPlain.String())
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	webfinger := fedihelper.WebFinger{
		Aliases: []string{path.GenActor(m.logic.Domain())},
		Links: []fedihelper.Link{
			{
				Href: path.GenActor(m.logic.Domain()),
				Rel:  "self",
				Type: libhttp.MimeAppActivityJSON.String(),
			},
			{
				Href: path.GenActor(m.logic.Domain()),
				Rel:  "self",
				Type: libhttp.MimeAppActivityLDJSON.String(),
			},
		},
		Subject: subject,
	}

	w.Header().Set("Content-Type", libhttp.MimeAppJSON.String())
	err := json.NewEncoder(w).Encode(webfinger)
	if err != nil {
		l.Errorf("marshaling json: %s", err.Error())
	}
}
