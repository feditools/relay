package models

// Actor represents an activity pub actor
type Actor struct {
	Context           string    `json:"@context"`
	Endpoints         Endpoints `json:"endpoints"`
	Followers         string    `json:"followers"`
	Following         string    `json:"following"`
	Inbox             string    `json:"inbox"`
	Name              string    `json:"name"`
	Type              string    `json:"type"`
	ID                string    `json:"id"`
	PublicKey         PublicKey `json:"publicKey"`
	Summary           string    `json:"summary"`
	PreferredUsername string    `json:"preferredUsername"`
	URL               string    `json:"url"`
}

// Endpoints represents known activity pub endpoints
type Endpoints struct {
	SharedInbox string `json:"sharedInbox"`
}

// PublicKey represents an actor's public key
type PublicKey struct {
	ID           string `json:"id"`
	Owner        string `json:"owner"`
	PublicKeyPEM string `json:"publicKeyPem"`
}
