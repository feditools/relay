package activitypub

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	apmodels "github.com/feditools/relay/internal/activitypub/models"
	"github.com/feditools/relay/internal/http"
	"github.com/feditools/relay/internal/models"
	"github.com/go-fed/httpsig"
	"github.com/tyrm/go-util/mimetype"
	nethttp "net/http"
	"net/url"
	"strings"
)

func extractPublicKeyFromActor(actor *apmodels.Actor) (*rsa.PublicKey, error) {
	l := logger.WithField("func", "extractPublicKeyFromActor")

	actorPem, _ := pem.Decode([]byte(actor.PublicKey.PublicKeyPEM))
	if actorPem == nil {
		msg := fmt.Sprintf("actor %s has invalid public key", actor.URL)
		l.Debugf(msg)
		return nil, errors.New(msg)
	}
	if actorPem.Type != "BEGIN PUBLIC KEY" {
		msg := fmt.Sprintf("actor %s has wrong key type", actor.URL)
		l.Debugf(msg)
		return nil, errors.New(msg)
	}

	parsedKey, err := x509.ParsePKIXPublicKey(actorPem.Bytes)
	if err != nil {
		msg := fmt.Sprintf("can't parse public key for %s", actor.URL)
		l.Debugf(msg)
		return nil, errors.New(msg)
	}
	pubKey, ok := parsedKey.(*rsa.PublicKey)
	if !ok {
		msg := fmt.Sprintf("can't cast public key for %s", actor.URL)
		l.Debugf(msg)
		return nil, errors.New(msg)
	}
	return pubKey, nil
}

func (m *Module) fetchActor(ctx context.Context, a *url.URL) (*apmodels.Actor, error) {
	l := logger.WithField("func", "fetchActor")

	v, err, shared := m.outgoingRequestGroup.Do(fmt.Sprintf("fetchactor-%s", a.String()), func() (interface{}, error) {
		body, err := m.transport.InstanceGet(ctx, a, mimetype.ApplicationActivityJSON)
		if err != nil {
			l.Errorf("instance get %s: %s", a.String(), err.Error())
			return nil, err
		}

		var newActor *apmodels.Actor
		err = json.Unmarshal(body, newActor)
		if err != nil {
			l.Errorf("unmarshal json %s: %s", a.String(), err.Error())
			return nil, err
		}

		return newActor, err
	})

	if err != nil {
		l.Errorf("singleflight (shared: %v): %s", shared, err.Error())
		return nil, err
	}

	actor := v.(*apmodels.Actor)
	return actor, nil
}

func (m *Module) getInstanceWithPublicKey(ctx context.Context, actorURI *url.URL) (*models.Instance, error) {
	l := logger.WithField("func", "getInstanceWithPublicKey")

	instance, err := m.db.ReadInstanceByDomain(ctx, actorURI.Host)
	if err != nil {
		l.Errorf("db read: %s", err.Error())
		return nil, err
	}

	if instance == nil {
		l.Debugf("creating instance %s from actor", actorURI.Host)
		instance, err = m.makeInstanceFromActor(ctx, actorURI)
		if err != nil {
			l.Errorf("make actor: %s", err.Error())
			return nil, err
		}
		return instance, nil
	}

	if instance.PublicKey == nil {
		// fetch remote actorURI
		actor, err := m.fetchActor(ctx, actorURI)
		if err != nil {
			l.Errorf("fetch actor: %s", err.Error())
			return nil, err
		}

		// make public key
		pubKey, err := extractPublicKeyFromActor(actor)
		if err != nil {
			l.Errorf("extracting public key: %s", err.Error())
			return nil, err
		}

		instance.PublicKey = pubKey
	}

	return nil, nil
}

func (m *Module) makeInstanceFromActor(ctx context.Context, actorURI *url.URL) (*models.Instance, error) {
	l := logger.WithField("func", "makeInstanceFromActor")

	// fetch remote actor
	actor, err := m.fetchActor(ctx, actorURI)
	if err != nil {
		l.Errorf("fetch actor: %s", err.Error())
		return nil, err
	}

	// make public key
	pubKey, err := extractPublicKeyFromActor(actor)
	if err != nil {
		l.Errorf("extracting public key: %s", err.Error())
		return nil, err
	}

	// create new instance
	newInstance := &models.Instance{
		Domain:    actorURI.Host,
		InboxIRI:  actor.Endpoints.SharedInbox,
		PublicKey: pubKey,
	}
	err = m.db.CreateInstance(ctx, newInstance)
	if err != nil {
		l.Errorf("db create: %s", err.Error())
		return nil, err
	}

	return newInstance, nil
}

func (m *Module) validateRequest(r *nethttp.Request, actorURI *url.URL) (bool, *models.Instance) {
	l := logger.WithField("func", "validateRequest")

	ctx := r.Context()

	// get verifier from context
	cVerifier := ctx.Value(http.ContextKeyKeyVerifier)
	if cVerifier == nil {
		l.Debug("verifier missing in context")
		return false, nil
	}
	verifier, ok := cVerifier.(httpsig.Verifier)
	if !ok {
		l.Warnf("can't cast verifier")
		return false, nil
	}

	// get signature from context
	/*signaturei := ctx.Value(http.ContextKeyHTTPSignature)
	if signaturei == nil {
		l.Debug("signature missing in context")
		return false
	}
	signature, ok := signaturei.(string)
	if !ok {
		l.Debug("couldn't extract signature")
		return false
	}*/

	// parse key uri
	publicKeyID, err := url.Parse(verifier.KeyId())
	if err != nil {
		l.Debug("can't parse public key URI")
		return false, nil
	}

	// relay should never talk to itself
	if strings.EqualFold(publicKeyID.Host, m.domain) {
		l.Warnf("received request from self")
		return false, nil
	}

	// get instance from database
	instance, err := m.getInstanceWithPublicKey(ctx, actorURI)
	if err != nil {
		l.Errorf("geting instance: %s", err.Error())
		return false, nil
	}

	// validate signature
	if instance.PublicKey == nil {
		// fetch remote actor
		actor, err := m.fetchActor(ctx, actorURI)
		if err != nil {
			l.Errorf("fetch actor: %s", err.Error())
			return false, nil
		}

		// make public key
		pubKey, err := extractPublicKeyFromActor(actor)
		if err != nil {
			l.Errorf("extracting public key: %s", err.Error())
			return false, nil
		}

		instance.PublicKey = pubKey
	}

	// try to verify known algos
	for _, algo := range validAlgs {
		err := verifier.Verify(instance.PublicKey, algo)
		if err == nil {
			l.Tracef("request passed %s algo", algo)
			return true, instance
		}
	}

	return false, nil
}
