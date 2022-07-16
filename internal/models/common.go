package models

// Link represents a link
type Link struct {
	Href     string `json:"href,omitempty"`
	Rel      string `json:"rel,omitempty"`
	Template string `json:"template,omitempty"`
	Type     string `json:"type,omitempty"`
}
