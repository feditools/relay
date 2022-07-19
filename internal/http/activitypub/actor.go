package activitypub

import (
	"encoding/json"
	"github.com/feditools/go-lib/fedihelper"
	libhttp "github.com/feditools/go-lib/http"
	"github.com/feditools/relay/internal/path"
	"net/http"
)

func (m *Module) actorGetHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "actorGetHandler")

	actor := m.genRelayActor()

	w.Header().Set("Content-Type", libhttp.MimeAppJSON.String())
	err := json.NewEncoder(w).Encode(actor)
	if err != nil {
		l.Errorf("marshaling json: %s", err.Error())
	}
}

func (m *Module) genRelayActor() *fedihelper.Actor {
	return &fedihelper.Actor{
		Context: ContextActivityStreams,
		Endpoints: fedihelper.Endpoints{
			SharedInbox: path.GenInbox(m.logic.Domain()),
		},
		Followers: path.GenFollowers(m.logic.Domain()),
		Following: path.GenFollowing(m.logic.Domain()),
		Inbox:     path.GenInbox(m.logic.Domain()),
		Name:      m.appName,
		Type:      "Application",
		ID:        path.GenActor(m.logic.Domain()),
		PublicKey: fedihelper.PublicKey{
			ID:           path.GenPublicKey(m.logic.Domain()),
			Owner:        path.GenActor(m.logic.Domain()),
			PublicKeyPEM: m.publicKeyPem,
		},
		Summary:           ActorSummary,
		PreferredUsername: "relay",
		URL:               path.GenActor(m.logic.Domain()),
	}
}
