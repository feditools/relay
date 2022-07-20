package webapp

import (
	"fmt"
	"github.com/feditools/go-lib/fedihelper"
	"github.com/feditools/relay/internal/models"
	"github.com/feditools/relay/internal/path"
	"github.com/feditools/relay/internal/token"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"net/http"
)

// CallbackOauthGetHandler handles an oauth callback.
func (m *Module) CallbackOauthGetHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "CallbackMastodonGetHandler")

	// lookup instance
	vars := mux.Vars(r)
	kind, id, err := m.tokz.DecodeToken(vars[path.VarInstanceID])
	if err != nil {
		l.Debugf("decode token: %s", err.Error())
		m.returnErrorPage(w, r, http.StatusBadRequest, "bad token")

		return
	}
	if kind != token.KindInstance {
		l.Debug("token is wrong kind")
		m.returnErrorPage(w, r, http.StatusBadRequest, "bad token")

		return
	}
	instance, err := m.db.ReadInstance(r.Context(), id)
	if err != nil {
		l.Errorf("db read instance: %s", err.Error())
		m.returnErrorPage(w, r, http.StatusInternalServerError, err.Error())

		return
	}
	if instance == nil {
		m.returnErrorPage(w, r, http.StatusNotFound, vars["token"])

		return
	}

	switch fedihelper.SoftwareName(instance.Software) {
	case fedihelper.SoftwareMastodon:
		// get code
		var code []string
		var ok bool
		if code, ok = r.URL.Query()["code"]; !ok || len(code[0]) < 1 {
			l.Debugf("missing code")
			m.returnErrorPage(w, r, http.StatusBadRequest, "missing code")

			return
		}

		// retrieve access token
		var accessToken string
		accessToken, err = m.fedi.Helper(fedihelper.SoftwareMastodon).GetAccessToken(
			r.Context(),
			path.GenCallbackOauth(m.domain, m.tokz.GetToken(instance)),
			instance,
			code[0],
		)
		if err != nil {
			l.Errorf("get access token error: %s", err.Error())
			m.returnErrorPage(w, r, http.StatusInternalServerError, err.Error())

			return
		}
		l.Debugf("access token: %s", accessToken)

		// retrieve current account
		accountI, err := m.fedi.Helper(fedihelper.SoftwareMastodon).GetCurrentAccount(
			r.Context(),
			instance,
			accessToken,
		)
		if err != nil {
			l.Errorf("get access token error: %s", err.Error())
			m.returnErrorPage(w, r, http.StatusInternalServerError, err.Error())

			return
		}
		account, ok := accountI.(*models.Account)
		if !ok {
			msg := "can't cast account to FediAccount"
			l.Error(msg)
			m.returnErrorPage(w, r, http.StatusInternalServerError, msg)

			return
		}

		// increment login
		err = m.db.IncAccountLoginCount(r.Context(), account)
		if err != nil {
			l.Errorf("db inc login: %s", err.Error())
			m.returnErrorPage(w, r, http.StatusInternalServerError, err.Error())

			return
		}

		// init session
		us := r.Context().Value(ContextKeySession).(*sessions.Session) // nolint
		us.Values[SessionKeyAccountID] = account.ID
		err = us.Save(r, w)
		if err != nil {
			m.returnErrorPage(w, r, http.StatusInternalServerError, err.Error())

			return
		}

		l.Debugf("account: %#v", account)

		// redirect to last page
		val := us.Values[SessionKeyLoginRedirect]
		var loginRedirect string
		if loginRedirect, ok = val.(string); !ok {
			// redirect home page if no login-redirect
			http.Redirect(w, r, path.AppHome, http.StatusFound)

			return
		}

		// Set login redirect to nil
		us.Values[SessionKeyLoginRedirect] = nil
		err = us.Save(r, w)
		if err != nil {
			m.returnErrorPage(w, r, http.StatusInternalServerError, err.Error())

			return
		}
		http.Redirect(w, r, loginRedirect, http.StatusFound)

		return
	default:
		m.returnErrorPage(w, r, http.StatusNotImplemented, fmt.Sprintf("no helper for '%s'", instance.Software))

		return
	}
}
