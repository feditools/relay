package models

// WebFinger is a web finger response.
type WebFinger struct {
	Subject string   `json:"subject"`
	Aliases []string `json:"aliases"`
	Links   []Link   `json:"links"`
}
