package models

// WebFinger represents a web finger response
type WebFinger struct {
	Aliases []string `json:"aliases,omitempty"`
	Links   []Link   `json:"links,omitempty"`
	Subject string   `json:"subject,omitempty"`
}
