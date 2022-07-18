package fedihelper

// Software is a federated social software type.
type Software string

const (
	HostMetaWebFingerTemplateRel = "lrdd"
	// NodeInfo20Schema the schema url for nodeinfo 2.0.
	NodeInfo20Schema = "http://nodeinfo.diaspora.software/ns/schema/2.0"
	// SoftwareMastodon is the software keyword for Mastodon.
	SoftwareMastodon Software = "mastodon"
)
