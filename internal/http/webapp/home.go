package webapp

import (
	"github.com/feditools/go-lib/language"
	"github.com/feditools/relay/internal/http/template"
	"github.com/feditools/relay/internal/path"
	"net/http"
)

// HomeGetHandler serves the home page.
func (m *Module) HomeGetHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "HomeGetHandler")

	// get localizer
	localizer := r.Context().Value(ContextKeyLocalizer).(*language.Localizer) //nolint

	// Init template variables
	tmplVars := &template.Home{}
	err := m.initTemplatePublic(w, r, tmplVars)
	if err != nil {
		m.returnErrorPage(w, r, http.StatusInternalServerError, err.Error())

		return
	}

	tmplVars.PageTitle = localizer.TextRelay(1).String()
	tmplVars.InboxHref = path.GenInbox(m.domain)
	tmplVars.ActorHref = path.GenActor(m.domain)

	followingInstance, err := m.db.ReadInstancesWhereFollowing(r.Context())
	if err != nil {
		l.Errorf("db read instances: %s", err.Error())
		m.returnErrorPage(w, r, http.StatusInternalServerError, ErrorResponseDBError)

		return
	}
	tmplVars.FollowingInstances = followingInstance

	blocks, err := m.db.ReadBlocks(r.Context())
	if err != nil {
		l.Errorf("db read blocks: %s", err.Error())
		m.returnErrorPage(w, r, http.StatusInternalServerError, ErrorResponseDBError)

		return
	}
	tmplVars.Blocks = blocks

	err = m.executeTemplate(w, template.HomeName, tmplVars)
	if err != nil {
		l.Errorf("could not render '%s' template: %s", template.HomeName, err.Error())
	}
}

// ForwardToHomeHandler serves a home forwarder.
func (m *Module) ForwardToHomeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, path.AppHome, http.StatusFound)
}