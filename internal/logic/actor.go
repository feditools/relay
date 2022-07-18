package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/feditools/relay/internal/models"
	"github.com/tyrm/go-util/mimetype"
	"net/url"
)

func (l *Logic) fetchActor(ctx context.Context, actorIRI *url.URL) (*models.Actor, error) {
	log := logger.WithField("func", "fetchActor")

	// do request
	v, err, shared := l.outgoingRequestGroup.Do(fmt.Sprintf("fetchactor-%s", actorIRI.String()), func() (interface{}, error) {
		// check cache
		cachedActor, ok := l.cacheActor.Get(actorIRI.String())
		if ok {
			return cachedActor, nil
		}

		// get actor data
		body, err := l.transport.InstanceGet(ctx, actorIRI, mimetype.ApplicationActivityJSON)
		if err != nil {
			log.Errorf("instance get %s: %s", actorIRI.String(), err.Error())
			return nil, err
		}

		// unmarshal json to object
		var newActor models.Actor
		err = json.Unmarshal(body, &newActor)
		if err != nil {
			log.Errorf("unmarshal json %s: %s", actorIRI.String(), err.Error())
			return nil, err
		}

		// update cache
		_ = l.cacheActor.Add(actorIRI.String(), newActor)

		return newActor, err
	})

	if err != nil {
		log.Errorf("singleflight (shared: %v): %s", shared, err.Error())
		return nil, err
	}

	actor := v.(models.Actor)
	return &actor, nil
}

func (l *Logic) getActorFromDomain(ctx context.Context, domain string) (*models.Actor, error) {
	log := logger.WithField("func", "getActorFromDomain")

	// pull host meta
	hostMeta, err := l.getHostMeta(ctx, domain)
	if err != nil {
		log.Debugf("can't retrieve host meta: %s", err.Error())

		return nil, NewErrorf("host meta: %s", err.Error())
	}

	// perform web finger
	webfingerURI := hostMeta.WebfingerURI()
	if webfingerURI == "" {
		log.Debug("host meta missing webfinger URI")

		return nil, NewError("host meta missing webfinger URI")
	}
	webFinger, err := l.fetchWebFinger(ctx, webfingerURI, domain, domain)
	if err != nil {
		log.Debugf("can't retrieve webfinger: %s", err.Error())

		return nil, NewErrorf("host meta: %s", err.Error())
	}

	// fetch actor
	actorIRI, err := webFinger.ActorURI()
	if err != nil {
		log.Debugf("can't get actor uri: %s", err.Error())

		return nil, NewErrorf("host meta: %s", err.Error())
	}
	actor, err := l.fetchActor(ctx, actorIRI)
	if err != nil {
		log.Debugf("can't fetch actor: %s", err.Error())

		return nil, NewErrorf("host meta: %s", err.Error())
	}

	return actor, nil
}
