package activitypub

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"github.com/feditools/relay/internal/models"
	"github.com/feditools/relay/internal/path"
)

func (m *Module) createInstanceSelf(ctx context.Context) (*models.Instance, error) {
	l := logger.WithField("func", "createInstanceSelf")

	// generate key
	privatekey, err := rsa.GenerateKey(rand.Reader, ActorKeySize)
	if err != nil {
		l.Errorf("genrating private key: %s", err.Error())
		return nil, err
	}

	publickey := &privatekey.PublicKey

	// create new instance
	newInstance := &models.Instance{
		Domain:   m.domain,
		ActorIRI: m.genActorSelf(),
		InboxIRI: path.GenInbox(m.domain),

		PublicKey:  publickey,
		PrivateKey: privatekey,
	}

	// add to database
	err = m.db.CreateInstance(ctx, newInstance)
	if err != nil {
		l.Errorf("db create: %s", err.Error())
		return nil, err
	}

	return newInstance, nil
}
