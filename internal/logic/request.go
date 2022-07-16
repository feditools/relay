package logic

import (
	"github.com/feditools/relay/internal/http"
	"github.com/feditools/relay/internal/models"
	"github.com/go-fed/httpsig"
	nethttp "net/http"
	"net/url"
	"strings"
)

func (m *Logic) ValidateRequest(r *nethttp.Request, actorURI *url.URL) (bool, *models.Instance) {
	log := logger.WithField("func", "validateRequest")

	ctx := r.Context()

	// get verifier from context
	cVerifier := ctx.Value(http.ContextKeyKeyVerifier)
	if cVerifier == nil {
		log.Debug("verifier missing in context")
		return false, nil
	}
	verifier, ok := cVerifier.(httpsig.Verifier)
	if !ok {
		log.Warnf("can't cast verifier")
		return false, nil
	}

	// get signature from context
	/*signaturei := ctx.Value(http.ContextKeyHTTPSignature)
	if signaturei == nil {
		log.Debug("signature missing in context")
		return false
	}
	signature, ok := signaturei.(string)
	if !ok {
		log.Debug("couldn't extract signature")
		return false
	}*/

	// parse key uri
	publicKeyID, err := url.Parse(verifier.KeyId())
	if err != nil {
		log.Debug("can't parse public key URI")
		return false, nil
	}

	// relay should never talk to itself
	if strings.EqualFold(publicKeyID.Host, m.domain) {
		log.Warnf("received request from self")
		return false, nil
	}

	// get instance from database
	instance, err := m.getInstanceWithPublicKey(ctx, actorURI)
	if err != nil {
		log.Errorf("geting instance: %s", err.Error())
		return false, nil
	}

	// validate signature
	if instance.PublicKey == nil {
		// fetch remote actor
		actor, err := m.fetchActor(ctx, actorURI)
		if err != nil {
			log.Errorf("fetch actor: %s", err.Error())
			return false, nil
		}

		// make public key
		pubKey, err := actor.RSAPublicKey()
		if err != nil {
			log.Errorf("extracting public key: %s", err.Error())
			return false, nil
		}

		instance.PublicKey = pubKey
	}

	// try to verify known algos
	for _, algo := range m.validAlgs {
		err := verifier.Verify(instance.PublicKey, algo)
		if err == nil {
			log.Tracef("request passed %s algo", algo)
			return true, instance
		}
	}

	return false, nil
}
