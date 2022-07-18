package activitypub

import (
	"context"
	"github.com/feditools/relay/internal/http"
	"github.com/go-fed/httpsig"
	"github.com/tyrm/go-util/mimetype"
	nethttp "net/http"
	"net/url"
)

type CTXVerifier struct {
	Verifier httpsig.Verifier
}

func (m *Module) middlewareCheckHTTPSig(next nethttp.Handler) nethttp.Handler {
	return nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		l := logger.WithField("func", "middlewareCheckHTTPSig")

		l.Debugf("running middleware")

		// create verifier
		verifier, err := httpsig.NewVerifier(r)
		if err != nil {
			l.Debugf("verifier error: %s", err.Error())
			next.ServeHTTP(w, r)
			return
		}

		// check key-id
		KeyIDURI, err := url.Parse(verifier.KeyId())
		if err != nil {
			l.Debugf("url parse error: %s", err.Error())
			next.ServeHTTP(w, r)
			return
		}
		if KeyIDURI == nil {
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), http.ContextKeyKeyVerifier, verifier)

		// check for domain block
		isBlocked, err := m.logic.IsDomainBlocked(ctx, KeyIDURI.Host)
		if err != nil {
			l.Errorf("is domain blocked: %s", err.Error())
			w.Header().Set("Content-Type", mimetype.TextPlain)
			nethttp.Error(w, nethttp.StatusText(nethttp.StatusInternalServerError), nethttp.StatusInternalServerError)
			return
		}
		if isBlocked {
			l.Debugf("domain %s is blocked", KeyIDURI.Host)
			w.Header().Set("Content-Type", mimetype.TextPlain)
			nethttp.Error(w, nethttp.StatusText(nethttp.StatusUnauthorized), nethttp.StatusUnauthorized)
			return
		}

		// get signature
		signature := r.Header.Get("signature")
		if signature != "" {
			ctx = context.WithValue(ctx, http.ContextKeyHTTPSignature, signature)
		}

		// do request with verifier
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
