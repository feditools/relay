package fedihelper

// Link represents a link
type Link struct {
	Href     string `json:"href,omitempty"`
	Rel      string `json:"rel,omitempty" xml:"rel,attr"`
	Template string `json:"template,omitempty" xml:"template,attr"`
	Type     string `json:"type,omitempty"`
}
