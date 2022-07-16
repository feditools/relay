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
			SharedInbox: path.GenInbox(m.logic.Domain()),
		},
		Followers: path.GenFollowers(m.logic.Domain()),
		Following: path.GenFollowing(m.logic.Domain()),
		Inbox:     path.GenInbox(m.logic.Domain()),
		Name:      m.appName,
		Type:      "Application",
		ID:        path.GenActor(m.logic.Domain()),
		PublicKey: apmodels.PublicKey{
			ID:           path.GenPublicKey(m.logic.Domain()),
			Owner:        path.GenActor(m.logic.Domain()),
			PublicKeyPEM: m.publicKeyPem,
		},
		Summary:           ActorSummary,
		PreferredUsername: "relay",
		URL:               path.GenActor(m.logic.Domain()),
	}
}
