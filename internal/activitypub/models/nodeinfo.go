package models

// NodeInfoWellKnown is a well known federated nodeinfo object
type NodeInfoWellKnown struct {
	Links []Link `json:"links,omitempty"`
}

// NodeInfo is a federated nodeinfo object
type NodeInfo struct {
	Metadata          map[string]interface{} `json:"metadata"`
	OpenRegistrations bool                   `json:"openRegistrations"`
	Protocols         []string               `json:"protocols"`
	Services          Services               `json:"services"`
	Software          Software               `json:"software"`
	Usage             Usage                  `json:"usage"`
	Version           string                 `json:"version"`
}

// Link represents a link
type Link struct {
	Href     string `json:"href,omitempty"`
	Rel      string `json:"rel,omitempty"`
	Template string `json:"template,omitempty"`
	Type     string `json:"type,omitempty"`
}

// Software contains the software and version of the node
type Software struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// Services contains the supported services of the node
type Services struct {
	Inbound  []string `json:"inbound"`
	Outbound []string `json:"outbound"`
}

// Usage contains usage statistics
type Usage struct {
	LocalPosts int64      `json:"localPosts"`
	Users      UsageUsers `json:"users"`
}

// UsageUsers contains usage statistics about users
type UsageUsers struct {
	Total int64 `json:"total"`
}
