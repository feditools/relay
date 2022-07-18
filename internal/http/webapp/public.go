package webapp

import (
	"github.com/feditools/go-lib/language"
	"github.com/feditools/relay/internal/http/template"
	"github.com/feditools/relay/internal/path"
	"net/http"
)

func makePublicNavbar(r *http.Request) template.Navbar {
	// get localizer
	localizer := r.Context().Value(ContextKeyLocalizer).(*language.Localizer) //nolint

	// create navbar
	newNavbar := template.Navbar{
		{
			Text:     localizer.TextHomeWeb().String(),
			MatchStr: path.ReHome,
			FAIcon:   "home",
			URL:      path.Home,
		},
	}

	newNavbar.ActivateFromPath(r.URL.Path)

	return newNavbar
}
