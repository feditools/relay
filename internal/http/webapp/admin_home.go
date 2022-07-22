package webapp

import (
	"github.com/feditools/go-lib/language"
	"github.com/feditools/relay/internal/http/template"
	"net/http"
)

// AdminHomeGetHandler serves the home page.
func (m *Module) AdminHomeGetHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "AdminHomeGetHandler")

	// get localizer
	localizer := r.Context().Value(ContextKeyLocalizer).(*language.Localizer) //nolint

	// Init template variables
	tmplVars := &template.AdminHome{
		Common: template.Common{
			PageTitle: localizer.TextRelay(1).String(),
		},
	}
	err := m.initTemplateAdmin(w, r, tmplVars)
	if err != nil {
		m.returnErrorPage(w, r, http.StatusInternalServerError, err.Error())

		return
	}

	err = m.executeTemplate(w, template.AdminHomeName, tmplVars)
	if err != nil {
		l.Errorf("could not render '%s' template: %s", template.AdminHomeName, err.Error())
	}
}
