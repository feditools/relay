package template

import (
	"github.com/feditools/go-lib/language"
	libtemplate "github.com/feditools/go-lib/template"
	"github.com/feditools/relay/internal/models"
)

// Common contains the variables used in nearly every template.
type Common struct {
	Language  string
	Localizer *language.Localizer

	Account *models.Account

	Alerts            *[]libtemplate.Alert
	FooterScripts     []libtemplate.Script
	FooterExtraScript string
	HeadLinks         []libtemplate.HeadLink
	AdminLink         string
	LoginLink         string
	LogoSrcDark       string
	LogoSrcLight      string
	LogoutLink        string
	NavBar            Navbar
	NavBarDark        bool
	PageTitle         string
}

// AddHeadLink adds a headder link to the template.
func (t *Common) AddHeadLink(l libtemplate.HeadLink) {
	if t.HeadLinks == nil {
		t.HeadLinks = []libtemplate.HeadLink{}
	}
	t.HeadLinks = append(t.HeadLinks, l)
}

// AddFooterScript adds a footer script to the template.
func (t *Common) AddFooterScript(s libtemplate.Script) {
	if t.FooterScripts == nil {
		t.FooterScripts = []libtemplate.Script{}
	}
	t.FooterScripts = append(t.FooterScripts, s)
}

// SetLanguage sets the template's default language.
func (t *Common) SetLanguage(l string) {
	t.Language = l
}

// SetLocalizer sets the localizer the template will use to generate text.
func (t *Common) SetLocalizer(l *language.Localizer) {
	t.Localizer = l
}

// SetLinks sets the template's links.
func (t *Common) SetLinks(admin, login, logout string) {
	t.AdminLink = admin
	t.LoginLink = login
	t.LogoutLink = logout
}

// SetLogoSrc sets the src for the logo image.
func (t *Common) SetLogoSrc(dark, light string) {
	t.LogoSrcDark = dark
	t.LogoSrcLight = light
}

// SetNavbar sets the top level navbar used by the template.
func (t *Common) SetNavbar(nodes Navbar) {
	t.NavBar = nodes
}

// SetNavbarDark sets the navbar theme.
func (t *Common) SetNavbarDark(dark bool) {
	t.NavBarDark = dark
}

// SetAccount sets the currently logged-in account.
func (t *Common) SetAccount(account *models.Account) {
	t.Account = account
}
