package webapp

import (
	"context"
	libhttp "github.com/feditools/go-lib/http"
	"net/http"
)

// Middleware runs on every http request.
func (m *Module) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := logger.WithField("func", "middleware")

		// Init Session
		us, err := m.store.Get(r, "relay")
		if err != nil {
			l.Errorf("get session: %s", err.Error())
			m.returnErrorPage(w, r, http.StatusInternalServerError, err.Error())

			return
		}
		ctx := context.WithValue(r.Context(), ContextKeySession, us)

		// Retrieve our account and type-assert it
		val := us.Values[SessionKeyAccountID]
		if accountID, ok := val.(int64); ok {
			// read federated accounts
			account, err := m.db.ReadAccount(ctx, accountID)
			if err != nil {
				l.Errorf("db read: %s", err.Error())
				m.returnErrorPage(w, r, http.StatusInternalServerError, err.Error())

				return
			}

			if account != nil {
				// read federated instance
				instance, err := m.db.ReadInstance(ctx, account.InstanceID)
				if err != nil {
					l.Errorf("db read: %s", err.Error())
					m.returnErrorPage(w, r, http.StatusInternalServerError, err.Error())

					return
				}
				account.Instance = instance

				ctx = context.WithValue(ctx, ContextKeyAccount, account)
			}
		}

		// create localizer
		lang := r.FormValue("lang")
		accept := r.Header.Get("Accept-Language")
		localizer, err := m.language.NewLocalizer(lang, accept)
		if err != nil {
			l.Errorf("could get localizer: %s", err.Error())
			m.returnErrorPage(w, r, http.StatusInternalServerError, err.Error())

			return
		}
		ctx = context.WithValue(ctx, ContextKeyLocalizer, localizer)

		// set request language
		ctx = context.WithValue(ctx, ContextKeyLanguage, libhttp.GetPageLang(lang, accept, m.language.Language().String()))

		// Do Request
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
