package template

import (
	"regexp"

	libtemplate "github.com/feditools/go-lib/template"
)

// Navbar is a navbar that can be added to a page.
type Navbar []NavbarNode

// ActivateFromPath sets the active bool based on the match regex.
func (n *Navbar) ActivateFromPath(path string) {
	libtemplate.SetActive(n, path)
}

// GetChildren returns the children of the node or nil if no children.
func (n *Navbar) GetChildren(i int) libtemplate.ActivableSlice {
	if len((*n)[i].Children) == 0 {
		return nil
	}

	return &(*n)[i].Children
}

// GetMatcher returns the matcher of the node or nil if no matcher.
func (n *Navbar) GetMatcher(i int) *regexp.Regexp {
	return (*n)[i].MatchStr
}

// SetActive sets the active bool based on the match regex.
func (n *Navbar) SetActive(i int, a bool) {
	(*n)[i].Active = a
}

// Len returns the matcher of the node or nil if no matcher.
func (n *Navbar) Len() int {
	return len(*n)
}

// NavbarNode is an entry on a navbar, can be nested one level.
type NavbarNode struct {
	Text     string
	URL      string
	MatchStr *regexp.Regexp
	FAIcon   string

	Active   bool
	Disabled bool

	Children Navbar
}
