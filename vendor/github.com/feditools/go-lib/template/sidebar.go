package template

import "regexp"

// Sidebar is a sidebar that can be added to a page.
type Sidebar []SidebarNode

// ActivateFromPath sets the active bool based on the match regex.
func (s *Sidebar) ActivateFromPath(path string) {
	SetActive(s, path)
}

// GetChildren returns the children of the node or nil if no children.
func (s *Sidebar) GetChildren(i int) ActivableSlice {
	if len(*s) == 0 {
		return nil
	}

	return &(*s)[i].Children
}

// GetMatcher returns the matcher of the node or nil if no matcher.
func (s *Sidebar) GetMatcher(i int) *regexp.Regexp {
	return (*s)[i].Matcher
}

// SetActive sets the active bool based on the match regex.
func (s *Sidebar) SetActive(i int, a bool) {
	(*s)[i].Active = a
}

// Len returns the matcher of the node or nil if no matcher.
func (s *Sidebar) Len() int {
	return len(*s)
}

// SidebarNode is an entry on a sidebar.
type SidebarNode struct {
	Text    string
	URI     string
	Matcher *regexp.Regexp
	Icon    string
	Label   string

	Active   bool
	Disabled bool

	Children Sidebar
}
