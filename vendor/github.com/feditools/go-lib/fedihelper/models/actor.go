package models

// Actor is an actor response.
type Actor struct {
	Name              string `json:"name"`
	PreferredUsername string `json:"preferredUsername"`
	Type              string `json:"type"`
}
