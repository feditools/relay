package webapp

import (
	"github.com/feditools/go-lib/language"
	"github.com/feditools/relay/internal/http/template"
	"github.com/feditools/relay/internal/path"
	nethttp "net/http"
)

func makeAdminNavbar(r *nethttp.Request) template.Navbar {
	// get localizer
	l := r.Context().Value(ContextKeyLocalizer).(*language.Localizer) // nolint

	// create navbar
	newNavbar := template.Navbar{
		{
			Text:     l.TextModeration().String(),
			MatchStr: path.ReAppAdminHome,
			FAIcon:   "arrows-turn-to-dots",
			URL:      path.AppAdminHome,
		},
	}

	newNavbar.ActivateFromPath(r.URL.Path)

	return newNavbar
}
