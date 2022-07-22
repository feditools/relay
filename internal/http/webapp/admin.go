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
			Text:     l.TextInstance(2).String(),
			MatchStr: path.ReAppAdminInstancesPre,
			FAIcon:   "server",
			URL:      path.AppAdminInstanceList,
		},
		{
			Text:     l.TextBlock(2).String(),
			MatchStr: path.ReAppAdminBlockPre,
			FAIcon:   "user-slash",
			URL:      path.AppAdminBlockList,
		},
	}

	newNavbar.ActivateFromPath(r.URL.Path)

	return newNavbar
}
