package activitypub

import (
	"encoding/json"
	apmodels "github.com/feditools/relay/internal/activitypub/models"
	"github.com/feditools/relay/internal/path"
	"github.com/tyrm/go-util/mimetype"
	"net/http"
)

func (m *Module) actorGetHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "actorGetHandler")

	actor := m.genRelayActor()

	w.Header().Set("Content-Type", mimetype.ApplicationJSON)
	err := json.NewEncoder(w).Encode(actor)
	if err != nil {
		l.Errorf("marshaling json: %s", err.Error())
	}
}

func (m *Module) genRelayActor() *apmodels.Actor {
	return &apmodels.Actor{
		Context: ContextActivityStreams,
		Endpoints: apmodels.Endpoints{
			SharedInbox: path.GenInbox(m.domain),
		},
		Followers: path.GenFollowers(m.domain),
		Following: path.GenFollowing(m.domain),
		Inbox:     path.GenInbox(m.domain),
		Name:      m.appName,
		Type:      "Application",
		ID:        path.GenActor(m.domain),
		PublicKey: apmodels.PublicKey{
			ID:           path.GenPublicKey(m.domain),
			Owner:        path.GenActor(m.domain),
			PublicKeyPEM: m.publicKeyPem,
		},
		Summary:           ActorSummary,
		PreferredUsername: "relay",
		URL:               path.GenActor(m.domain),
	}
}
