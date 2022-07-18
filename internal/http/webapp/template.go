package webapp

import (
	"bytes"
	"github.com/feditools/go-lib/language"
	"github.com/feditools/relay/internal/http/template"
	"github.com/feditools/relay/internal/models"
	"net/http"
)

func (m *Module) executeTemplate(w http.ResponseWriter, name string, tmplVars interface{}) error {
	b := new(bytes.Buffer)
	err := m.templates.ExecuteTemplate(b, name, tmplVars)
	if err != nil {
		return err
	}

	if m.minify == nil {
		_, err := w.Write(b.Bytes())

		return err
	}

	return m.minify.Minify("text/html", w, b)
}

func (m *Module) initTemplate(_ http.ResponseWriter, r *http.Request, tmpl template.InitTemplate) error {
	// l := logger.WithField("func", "initTemplate")

	// set text handler
	localizer := r.Context().Value(ContextKeyLocalizer).(*language.Localizer) // nolint
	tmpl.SetLocalizer(localizer)

	// set language
	lang := r.Context().Value(ContextKeyLanguage).(string) // nolint
	tmpl.SetLanguage(lang)

	// set logo image src
	tmpl.SetLogoSrc(m.logoSrcDark, m.logoSrcLight)

	// add css
	for _, link := range m.headLinks {
		tmpl.AddHeadLink(link)
	}

	// add scripts
	for _, script := range m.footerScripts {
		tmpl.AddFooterScript(script)
	}

	if r.Context().Value(ContextKeyAccount) != nil {
		account := r.Context().Value(ContextKeyAccount).(*models.Account) // nolint
		tmpl.SetAccount(account)
	}

	// try to read session data
	/*if r.Context().Value(http.ContextKeySession) == nil {
		return nil
	}

	us := r.Context().Value(http.ContextKeySession).(*sessions.Session)
	saveSession := false

	if saveSession {
		err := us.Save(r, w)
		if err != nil {
			l.Warningf("initTemplate could not save session: %s", err.Error())
			return err
		}
	}*/

	return nil
}

func (m *Module) initTemplateAdmin(w http.ResponseWriter, r *http.Request, tmpl template.InitTemplate) error {
	err := m.initTemplate(w, r, tmpl)
	if err != nil {
		return err
	}

	// make admin navbar
	navbar := makeAdminNavbar(r)
	tmpl.SetNavbar(navbar)

	return nil
}

func (m *Module) initTemplatePublic(w http.ResponseWriter, r *http.Request, tmpl template.InitTemplate) error {
	err := m.initTemplate(w, r, tmpl)
	if err != nil {
		return err
	}

	// make admin navbar
	navbar := makePublicNavbar(r)
	tmpl.SetNavbar(navbar)

	return nil
}
