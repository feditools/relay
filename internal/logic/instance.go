package logic

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"github.com/feditools/relay/internal/config"
	"github.com/feditools/relay/internal/models"
	"github.com/feditools/relay/internal/path"
	"github.com/spf13/viper"
	"net/url"
	"time"
)

// GetPeers returns a list of peers
func (l *Logic) GetPeers(ctx context.Context) (*[]string, error) {
	log := logger.WithField("func", "GetPeers")

	// check cache
	peers, valid := l.getCachedPeerList()
	if valid {
		return peers, nil
	}

	// read from db
	instances, err := l.db.ReadInstancesWhereFollowing(ctx)
	if err != nil {
		log.Errorf("db read: %s", err.Error())
		return nil, err
	}
	log.Debugf("got %d instances", len(instances))

	// populate peer list
	newPeers := make([]string, len(instances))
	for i, instance := range instances {
		newPeers[i] = instance.Domain
	}

	// update cache
	l.setCachedPeerList(&newPeers)

	return &newPeers, nil
}

func (l *Logic) GetInstanceSelf(ctx context.Context) (*models.Instance, error) {
	log := logger.WithField("func", "GetInstanceSelf")

	instance := new(models.Instance)
	var err error
	instance, err = l.db.ReadInstanceByDomain(ctx, l.domain)
	if err != nil {
		log.Errorf("db read: %s", err.Error())

		return nil, err
	}

	return instance, nil
}

func (l *Logic) createInstanceSelf(ctx context.Context) (*models.Instance, error) {
	log := logger.WithField("func", "createInstanceSelf")

	// generate key
	privatekey, err := rsa.GenerateKey(rand.Reader, viper.GetInt(config.Keys.ActorKeySize))
	if err != nil {
		log.Errorf("genrating private key: %s", err.Error())
		return nil, err
	}

	publickey := &privatekey.PublicKey

	// create new instance
	newInstance := &models.Instance{
		Domain:   l.domain,
		ActorIRI: genActorSelf(l.domain),
		InboxIRI: path.GenInbox(l.domain),

		PublicKey:  publickey,
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

// peer list "cache" functions
func (l *Logic) getCachedPeerList() (*[]string, bool) {
	l.cPeerListLock.RLock()
	if l.cPeerListExpires.Before(time.Now()) {
		l.cPeerListLock.RUnlock()
		return nil, false
	}
	count := l.cPeerList
	l.cPeerListLock.RUnlock()
	return count, true
}

func (l *Logic) setCachedPeerList(peers *[]string) {
	l.cPeerListLock.Lock()
	l.cPeerList = peers
	l.cPeerListExpires = time.Now().Add(l.cPeerListValidity)
	l.cPeerListLock.Unlock()
}

func (l *Logic) getOrCreateSelfInstance(ctx context.Context) (*models.Instance, error) {
	log := logger.WithField("func", "getOrCreateSelfInstance")

	instance, err := l.GetInstanceSelf(ctx)
	if err != nil {
		log.Errorf("db read: %s", err.Error())

		return nil, err
	}

	if instance == nil {
		instance, err = l.createInstanceSelf(ctx)
		if err != nil {
			log.Errorf("create self: %s", err.Error())

			return nil, err
		}
	}

	return instance, nil
}

func (m *Logic) getInstanceWithPublicKey(ctx context.Context, actorURI *url.URL) (*models.Instance, error) {
	log := logger.WithField("func", "getInstanceWithPublicKey")

	instance, err := m.db.ReadInstanceByDomain(ctx, actorURI.Host)
	if err != nil {
		log.Errorf("db read: %s", err.Error())
		return nil, err
	}

	if instance == nil {
		log.Debugf("creating instance %s from actor", actorURI.Host)
		instance, err = m.makeInstanceFromActor(ctx, actorURI)
		if err != nil {
			log.Errorf("make actor: %s", err.Error())
			return nil, err
		}
		return instance, nil
	}

	if instance.PublicKey == nil {
		// fetch remote actorURI
		actor, err := m.fetchActor(ctx, actorURI)
		if err != nil {
			log.Errorf("fetch actor: %s", err.Error())
			return nil, err
		}

		// make public key
		pubKey, err := actor.RSAPublicKey()
		if err != nil {
			log.Errorf("extracting public key: %s", err.Error())
			return nil, err
		}

		instance.PublicKey = pubKey
	}

	return instance, nil
}

func (m *Logic) makeInstanceFromActor(ctx context.Context, actorURI *url.URL) (*models.Instance, error) {
	log := logger.WithField("func", "makeInstanceFromActor")

	// fetch remote actor
	actor, err := m.fetchActor(ctx, actorURI)
	if err != nil {
		log.Errorf("fetch actor: %s", err.Error())
		return nil, err
	}

	// make public key
	pubKey, err := actor.RSAPublicKey()
	if err != nil {
		log.Errorf("extracting public key: %s", err.Error())
		return nil, err
	}

	// create new instance
	newInstance := &models.Instance{
		Domain:    actorURI.Host,
		ActorIRI:  actorURI.String(),
		InboxIRI:  actor.Endpoints.SharedInbox,
		PublicKey: pubKey,
	}
	err = m.db.CreateInstance(ctx, newInstance)
	if err != nil {
		log.Errorf("db create: %s", err.Error())
		return nil, err
	}

	return newInstance, nil
}
