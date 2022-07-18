package logic

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"github.com/feditools/relay/internal/config"
	"github.com/feditools/relay/internal/db"
	"github.com/feditools/relay/internal/models"
	"github.com/feditools/relay/internal/path"
	"github.com/spf13/viper"
)

func (l *Logic) createInstanceSelf(ctx context.Context) (*models.Instance, error) {
	log := logger.WithField("func", "createInstanceSelf")

	// generate key
	privatekey, err := rsa.GenerateKey(rand.Reader, viper.GetInt(config.Keys.ActorKeySize))
	if err != nil {
		log.Errorf("genrating private key: %s", err.Error())
		return nil, err
	}

	// create new instance
	newInstance := &models.Instance{
		AccountDomain: l.domain,
		Domain:        l.domain,
		ActorIRI:      genActorSelf(l.domain),
		InboxIRI:      path.GenInbox(l.domain),

		PublicKey:  &privatekey.PublicKey,
		PrivateKey: privatekey,
	}

	// add to database
	err = l.db.CreateInstance(ctx, newInstance)
	if err != nil {
		log.Errorf("db create: %s", err.Error())
		return nil, err
	}

	return newInstance, nil
}

func (l *Logic) GetInstanceSelf(ctx context.Context) (*models.Instance, error) {
	log := logger.WithField("func", "GetInstanceSelf")

	instance := new(models.Instance)
	var err error
	instance, err = l.db.ReadInstanceByDomain(ctx, l.domain)
	if err != nil {
		if !errors.Is(err, db.ErrNoEntries) {
			log.Errorf("db read: %s", err.Error())
		}

		return nil, err
	}

	return instance, nil
}

func (l *Logic) getOrCreateSelfInstance(ctx context.Context) (*models.Instance, error) {
	log := logger.WithField("func", "getOrCreateSelfInstance")

	instance, err := l.GetInstanceSelf(ctx)
	if err == nil {
		return instance, nil
	}
	if err != nil && !errors.Is(err, db.ErrNoEntries) {
		log.Errorf("db read: %s", err.Error())

		return nil, err
	}

	newInstance, err := l.createInstanceSelf(ctx)
	if err != nil {
		log.Errorf("create self: %s", err.Error())

		return nil, err
	}

	return newInstance, nil
}
