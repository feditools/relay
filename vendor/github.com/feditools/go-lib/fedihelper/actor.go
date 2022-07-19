package fedihelper

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/feditools/go-lib/http"
	"net/url"
)

// Actor is an actor response.
type Actor struct {
	Context           interface{} `json:"@context"`
	Endpoints         Endpoints   `json:"endpoints"`
	Followers         string      `json:"followers"`
	Following         string      `json:"following"`
	Inbox             string      `json:"inbox"`
	Name              string      `json:"name"`
	Type              ActorType   `json:"type"`
	ID                string      `json:"id"`
	PublicKey         PublicKey   `json:"publicKey"`
	Summary           string      `json:"summary"`
	PreferredUsername string      `json:"preferredUsername"`
	URL               string      `json:"url"`
}

func (a *Actor) MakeInstance(instance Instance) error {
	if a.Type != TypeApplication {
		return NewErrorf("actor is not type %s", TypeApplication)
	}

	actorID, err := url.Parse(a.ID)
	if a.Type != TypeApplication {
		return NewErrorf("can't parse id: %s", err.Error())
	}

	instance.SetDomain(a.PreferredUsername)
	instance.SetServerHostname(actorID.Host)
	instance.SetActorURI(a.ID)
	instance.SetInboxURI(a.Endpoints.SharedInbox)

	return nil
}

func (a *Actor) RSAPublicKey() (*rsa.PublicKey, error) {
	l := logger.WithField("func", "RSAPublicKey")

	actorPem, _ := pem.Decode([]byte(a.PublicKey.PublicKeyPEM))
	if actorPem == nil {
		msg := fmt.Sprintf("actor %s has invalid public key", a.URL)
		l.Debugf(msg)

		return nil, errors.New(msg)
	}

	l.Debugf("public key type: %s", actorPem.Type)
	if actorPem.Type != "PUBLIC KEY" {
		msg := fmt.Sprintf("actor %s has wrong key type", a.URL)
		l.Debugf(msg)

		return nil, errors.New(msg)
	}

	parsedKey, err := x509.ParsePKIXPublicKey(actorPem.Bytes)
	if err != nil {
		msg := fmt.Sprintf("can't parse public key for %s", a.URL)
		l.Debugf(msg)

		return nil, errors.New(msg)
	}
	pubKey, ok := parsedKey.(*rsa.PublicKey)
	if !ok {
		msg := fmt.Sprintf("can't cast public key for %s", a.URL)
		l.Debugf(msg)

		return nil, errors.New(msg)
	}

	return pubKey, nil
}

func (f *FediHelper) FetchActor(ctx context.Context, actorIRI *url.URL) (*Actor, error) {
	log := logger.WithField("func", "fetchActor")

	// do request
	v, err, shared := f.requestGroup.Do(fmt.Sprintf("fetchactor-%s", actorIRI.String()), func() (interface{}, error) {
		// check cache
		cache, err := f.kv.GetActor(ctx, actorIRI.String())
		if err != nil && err.Error() != "nil" {
			fhErr := NewErrorf("redis get: %s", err.Error())
			log.Error(fhErr.Error())

			return nil, fhErr
		}
		if err == nil {
			return unmarshalActor(cache)
		}

		// get actor data
		bodyBytes, err := f.http.InstanceGet(ctx, actorIRI, http.MimeAppActivityJSON)
		if err != nil {
			log.Errorf("instance get %s: %s", actorIRI.String(), err.Error())
			return nil, err
		}

		// write cache
		err = f.kv.SetActor(ctx, actorIRI.String(), bodyBytes, f.actorCacheExp)
		if err != nil {
			fhErr := NewErrorf("redis set: %s", err.Error())
			log.Error(fhErr.Error())

			return nil, fhErr
		}

		return unmarshalActor(bodyBytes)
	})

	if err != nil {
		log.Errorf("singleflight (shared: %v): %s", shared, err.Error())
		return nil, err
	}

	actor := v.(*Actor)
	return actor, nil
}

func unmarshalActor(body []byte) (*Actor, error) {
	actor := new(Actor)
	if err := json.Unmarshal(body, actor); err != nil {
		return nil, NewErrorf("unmarshal: %s", err.Error())
	}

	return actor, nil
}
