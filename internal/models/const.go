package models

const (
	// ContextActivityStreams contains the context document for activity streams
	ContextActivityStreams = "https://www.w3.org/ns/activitystreams"

	HostMetaWebFingerTemplateRel = "lrdd"

	// MimeAppActivityJSON represents a JSON activity pub action type.
	MimeAppActivityJSON = `application/activity+json`

	// Activity Types

	// TypeAccept is the Accept activity Type
	TypeAccept = "Accept"
	// TypeAnnounce is the Announce activity Type
	TypeAnnounce = "Announce"
	// TypeCreate is the Create activity Type
	TypeCreate = "Create"
	// TypeDelete is the Delete activity Type
	TypeDelete = "Delete"
	// TypeFollow is the Follow activity Type
	TypeFollow = "Follow"
	// TypeUndo is the Undo activity Type
	TypeUndo = "Undo"
	// TypeUpdate is the Update activity Type
	TypeUpdate = "Update"

	// Actor Types

	// TypeApplication is the Application actor type
	TypeApplication = "Application"
	// TypePerson is the Person actor type
	TypePerson = "Person"
)
