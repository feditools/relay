package webapp

import (
	"github.com/feditools/relay/internal/path"
	"github.com/gorilla/sessions"
	"net/http"
)

// LogoutGetHandler logs a user out.
func (m *Module) LogoutGetHandler(w http.ResponseWriter, r *http.Request) {
	// Init Session
	us := r.Context().Value(ContextKeySession).(*sessions.Session) // nolint

	// Set account to nil
	us.Values[SessionKeyAccountID] = nil

	if err := us.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		m.returnErrorPage(w, r, http.StatusInternalServerError, err.Error())

		return
	}

	http.Redirect(w, r, path.AppHome, http.StatusFound)
}
