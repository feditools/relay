package webapp

import (
	"errors"
	"github.com/feditools/go-lib"
	"github.com/feditools/go-lib/language"
	libtemplate "github.com/feditools/go-lib/template"
	"github.com/feditools/relay/internal/db"
	"github.com/feditools/relay/internal/http/template"
	"github.com/feditools/relay/internal/models"
	"github.com/feditools/relay/internal/path"
	nethttp "net/http"
	"strings"
)

// LoginGetHandler serves the login page.
func (m *Module) LoginGetHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	m.displayLoginPage(w, r, "", "")
}

// LoginPostHandler attempts a login.
func (m *Module) LoginPostHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	l := logger.WithField("func", "LoginPostHandler")

	// parse form data
	if err := r.ParseForm(); err != nil {
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())

		return
	}

	// split account
	formAccount := r.Form.Get("account")
	_, domain, err := lib.SplitAccount(formAccount)
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusBadRequest, err.Error())

		return
	}

	// get instance
	var instance *models.Instance
	instance, err = m.db.ReadInstanceByDomain(r.Context(), domain)
	if err != nil && !errors.Is(err, db.ErrNoEntries) {
		// actual db error
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())

		return
	}
	if errors.Is(err, db.ErrNoEntries) {
		instance, err = m.fedi.GenInstanceFromDomain(r.Context(), domain)
		if err != nil {
			m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())

			return
		}

		err = m.db.CreateInstance(r.Context(), instance)
		if err != nil {
			m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())

			return
		}
	}

	// check if account exists
	loginURL, err := m.fedi.GetLoginURL(
		r.Context(),
		path.GenCallbackOauth(m.domain, m.tokz.GetToken(instance)),
		instance,
	)
	if err != nil {
		l.Errorf("get login url: %s", err.Error())
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())

		return
	}

	nethttp.Redirect(w, r, loginURL.String(), nethttp.StatusFound)
}

func (m *Module) displayLoginPage(w nethttp.ResponseWriter, r *nethttp.Request, account, formError string) {
	l := logger.WithField("func", "displayLoginPage")

	// get localizer
	localizer := r.Context().Value(ContextKeyLocalizer).(*language.Localizer) // nolint

	// Init template variables
	tmplVars := &template.Login{}
	err := m.initTemplate(w, r, tmplVars)
	if err != nil {
		nethttp.Error(w, err.Error(), nethttp.StatusInternalServerError)

		return
	}

	// add error css file
	signature, err := m.getSignatureCached(strings.TrimPrefix(path.FileLoginCSS, "/"))
	if err != nil {
		l.Errorf("getting signature for %s: %s", path.FileLoginCSS, err.Error())
	}
	tmplVars.AddHeadLink(libtemplate.HeadLink{
		HRef:        path.FileLoginCSS,
		Rel:         "stylesheet",
		CrossOrigin: "anonymous",
		Integrity:   signature,
	})

	tmplVars.PageTitle = localizer.TextLogin().String()

	// set form values
	tmplVars.FormError = formError
	tmplVars.FormAccount = account

	err = m.executeTemplate(w, template.LoginName, tmplVars)
	if err != nil {
		l.Errorf("could not render %s template: %s", template.LoginName, err.Error())
	}
}
