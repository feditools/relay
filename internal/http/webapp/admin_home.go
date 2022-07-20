package webapp

import (
	"github.com/feditools/go-lib/language"
	libtemplate "github.com/feditools/go-lib/template"
	"github.com/feditools/relay/internal/http/template"
	"github.com/feditools/relay/internal/path"
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
		Admin: template.Admin{
			Sidebar: makeAdminSidebar(r),
		},
	}
	err := m.initTemplateAdmin(w, r, tmplVars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	err = m.executeTemplate(w, template.AdminHomeName, tmplVars)
	if err != nil {
		l.Errorf("could not render '%s' template: %s", template.AdminHomeName, err.Error())
	}
}

func makeAdminSidebar(r *http.Request) libtemplate.Sidebar {
	// get localizer
	localizer := r.Context().Value(ContextKeyLocalizer).(*language.Localizer) // nolint

	// create sidebar
	newSidebar := libtemplate.Sidebar{
		{
			Children: []libtemplate.SidebarNode{
				{
					Text:    localizer.TextDashboard(1).String(),
					Matcher: path.ReAppAdminHome,
					Icon:    "desktop",
					URI:     path.AppAdminHome,
				},
			},
		},
		{
			Text: localizer.TextInstance(2).String(),
			Children: []libtemplate.SidebarNode{
				{
					Text:    localizer.TextInstance(2).String(),
					Matcher: path.ReAppAdminInstances,
					Icon:    "desktop",
					URI:     path.AppAdminInstances,
				},
			},
		},
	}

	newSidebar.ActivateFromPath(r.URL.Path)

	return newSidebar
}
