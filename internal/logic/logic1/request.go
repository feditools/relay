package logic1

import (
	"github.com/feditools/go-lib/fedihelper"
	"github.com/go-fed/httpsig"
	nethttp "net/http"
	"net/url"
	"strings"
)

func (l *Logic) ValidateRequest(r *nethttp.Request, actorURI *url.URL) (bool, *fedihelper.Actor) {
	log := logger.WithField("func", "ValidateRequest")

	ctx := r.Context()

	// create verifier
	verifier, err := httpsig.NewVerifier(r)
	if err != nil {
		log.Debugf("verifier error: %s", err.Error())

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

	// check for domain block
	isBlocked, err := l.IsDomainBlocked(ctx, publicKeyID.Host)
	if err != nil {
		log.Debugf("can't get domain block: %s", err.Error())

		return false, nil
	}
	if isBlocked {
		log.Debugf("domain %s is blocked", publicKeyID.Host)

		return false, nil
	}

	// fetch actor
	actor, err := l.fedi.FetchActor(ctx, actorURI)
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
