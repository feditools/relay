package activitypub

import "github.com/go-fed/httpsig"

const (

	// ActorKeySize is the key size used to generate a new key
	ActorKeySize = 2048
	// ActorSummary is the description of the actor for the relay
	ActorSummary = "Feditools ActivityPub Relay - https://github.com/feditools/relay"

	// ContextActivityStreams contains the context document for activity streams
	ContextActivityStreams = "https://www.w3.org/ns/activitystreams"

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
)

var (
	validAlgs = []httpsig.Algorithm{
		httpsig.RSA_SHA512,
		httpsig.RSA_SHA256,
		httpsig.ED25519,
	}
)
