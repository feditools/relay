package logic

import (
	"github.com/feditools/relay/internal/http"
	"github.com/feditools/relay/internal/models"
	"github.com/go-fed/httpsig"
	nethttp "net/http"
	"net/url"
	"strings"
)

func (l *Logic) ValidateRequest(r *nethttp.Request, actorURI *url.URL) (bool, *models.Actor) {
	log := logger.WithField("func", "ValidateRequest")

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

	// parse key uri
	publicKeyID, err := url.Parse(verifier.KeyId())
	if err != nil {
		log.Debug("can't parse public key URI")
		return false, nil
	}

	// relay should never talk to itself
	if strings.EqualFold(publicKeyID.Host, l.domain) {
		log.Warnf("received request from self")
		return false, nil
	}

	// TODO: check blocks

	// fetch actor
	actor, err := l.fetchActor(ctx, actorURI)
	if err != nil {
		log.Errorf("fetch actor: %s", err.Error())
		return false, nil
	}

	pk, err := actor.RSAPublicKey()
	if err != nil {
		log.Errorf("getting actor rsa public key: %s", err.Error())
		return false, nil
	}

	// try to verify known algos
	for _, algo := range l.validAlgs {
		err := verifier.Verify(pk, algo)
		if err == nil {
			log.Tracef("request passed %s algo", algo)
			return true, actor
		}
	}

	return false, nil
}
