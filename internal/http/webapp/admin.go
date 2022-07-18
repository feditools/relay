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
			Text:     l.TextHomeWeb().String(),
			MatchStr: path.ReAdmin,
			FAIcon:   "home",
			URL:      path.Admin,
		},
	}

	newNavbar.ActivateFromPath(r.URL.Path)

	return newNavbar
}
