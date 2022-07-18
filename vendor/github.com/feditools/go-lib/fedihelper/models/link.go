package models

// Link represents a link in an api response.
type Link struct {
	HRef string `json:"href"`
	Rel  string `json:"rel"`
	Type string `json:"type"`
}
