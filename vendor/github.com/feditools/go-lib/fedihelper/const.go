package fedihelper

// SoftwareName is a federated social software type.

type (
	ActivityType string
	ActorType    string
	SoftwareName string
)

func (a ActivityType) String() string {
	return string(a)
}

func (a ActorType) String() string {
	return string(a)
}

func (s SoftwareName) String() string {
	return string(s)
}

const (
	// ContextActivityStreams contains the context document for activity streams
	ContextActivityStreams = "https://www.w3.org/ns/activitystreams"
	// HostMetaWebFingerTemplateRel matches a webfinger link relationship.
	HostMetaWebFingerTemplateRel = "lrdd"
	// NodeInfo20Schema the schema url for nodeinfo 2.0.
	NodeInfo20Schema = "http://nodeinfo.diaspora.software/ns/schema/2.0"
	// SoftwareMastodon is the software keyword for Mastodon.
	SoftwareMastodon SoftwareName = "mastodon"

	// Activity Types

	// TypeAccept is the Accept activity Type
	TypeAccept ActivityType = "Accept"
	// TypeAnnounce is the Announce activity Type
	TypeAnnounce ActivityType = "Announce"
	// TypeCreate is the Create activity Type
	TypeCreate ActivityType = "Create"
	// TypeDelete is the Delete activity Type
	TypeDelete ActivityType = "Delete"
	// TypeFollow is the Follow activity Type
	TypeFollow ActivityType = "Follow"
	// TypeUndo is the Undo activity Type
	TypeUndo ActivityType = "Undo"
	// TypeUpdate is the Update activity Type
	TypeUpdate ActivityType = "Update"

	// Actor Types

	// TypeApplication is the Application actor type
	TypeApplication ActorType = "Application"
	// TypePerson is the Person actor type
	TypePerson ActorType = "Person"
)
