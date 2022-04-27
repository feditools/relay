package activitypub

import (
	"context"
	"fmt"
	rhttp "github.com/feditools/relay/internal/http"
	"github.com/go-fed/httpsig"
	"github.com/tyrm/go-util/mimetype"
	"net/http"
	"net/url"
)

func (m *Module) middlewareCheckHTTPSig(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := logger.WithField("func", "middlewareCheckHTTPSig")

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
		ctx := context.WithValue(r.Context(), rhttp.ContextKeyKeyVerifier, verifier)

		// check for domain block
		isBlocked, err := m.logic.IsDomainBlocked(ctx, KeyIDURI.Host)
		if err != nil {
			l.Errorf("is domain blocked: %s", err.Error())
			w.Header().Set("Content-Type", mimetype.TextPlain)
			w.Write([]byte(fmt.Sprintf("%d %s", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))))
			return
		}
		if isBlocked {
			l.Debugf("domain %s is blocked", KeyIDURI.Host)
			w.Header().Set("Content-Type", mimetype.TextPlain)
			w.Write([]byte(fmt.Sprintf("%d %s", http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))))
			return
		}

		// get signature
		signature := r.Header.Get("signature")
		if signature != "" {
			ctx = context.WithValue(ctx, rhttp.ContextKeyHTTPSignature, signature)
		}

		// do request with verifier
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
