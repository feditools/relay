package logic1

import (
	"context"
	"errors"
	"fmt"
	"github.com/feditools/relay/internal/db"
	"github.com/feditools/relay/internal/models"
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
		newPeers[i] = instance.ServerHostname
	}

	// update cache
	l.setCachedPeerList(&newPeers)

	return &newPeers, nil
}

func (l *Logic) GetInstance(ctx context.Context, domain string) (*models.Instance, error) {
	log := logger.WithField("func", "GetInstance")

	// try to get instance from db
	instance := new(models.Instance)
	var err error
	instance, err = l.db.ReadInstanceByDomain(ctx, domain)
	if err != nil {
		if !errors.Is(err, db.ErrNoEntries) {
			log.Errorf("db read: %s", err.Error())
		}

		return nil, err
	}

	return instance, nil
}

func (l *Logic) GetInstanceForActor(ctx context.Context, actorID *url.URL) (*models.Instance, error) {
	log := logger.WithField("func", "GetInstanceForActor")

	// try to get instance from db
	instance := new(models.Instance)
	var err error
	instance, err = l.db.ReadInstanceByActorIRI(ctx, actorID.String())
	if err == nil {
		return instance, nil
	}
	if err != nil && !errors.Is(err, db.ErrNoEntries) {
		log.Errorf("db read: %s", err.Error())

		return nil, err
	}

	// not in db, fetch actor
	actor, err := l.fedi.FetchActor(ctx, actorID)
	if err != nil {
		log.Errorf("fetching actor: %s", err.Error())

		return nil, fmt.Errorf("fetching actor: %s", err.Error())
	}
	newInstance := new(models.Instance)
	err = actor.MakeInstance(newInstance)
	if err != nil {
		log.Errorf("make instance: %s", err.Error())

		return nil, fmt.Errorf("make instance: %s", err.Error())
	}

	// add to db
	err = l.db.CreateInstance(ctx, newInstance)
	if err != nil {
		log.Errorf("db create: %s", err.Error())

		return nil, fmt.Errorf("db create: %s", err.Error())
	}

	return newInstance, nil
}

func (l *Logic) GetInstancesForForwarding(ctx context.Context, actorIRI, objectID string) ([]*models.Instance, error) {
	log := logger.WithField("func", "GetInstance")

	objectIDURI, err := url.Parse(objectID)
	if err != nil {
		log.Errorf("parsing object id uri: %s", err.Error())

		return nil, fmt.Errorf("parsing object id uri: %s", err.Error())
	}

	instances, err := l.db.ReadInstancesWhereFollowing(ctx)
	if err != nil {
		log.Errorf("db read: %s", err.Error())
		return nil, err
	}
	log.Debugf("got %d instances", len(instances))

	// remove sender and origin instances
	selectedInstances := make([]*models.Instance, 0)
	for _, instance := range instances {
		aIRI, err := url.Parse(instance.ActorIRI)
		if err != nil {
			log.Errorf("parsing object id uri: %s", err.Error())
			continue
		}

		if instance.ActorIRI != actorIRI && aIRI.Host != objectIDURI.Host {
			selectedInstances = append(selectedInstances, instance)
		}
	}

	return selectedInstances, nil
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
