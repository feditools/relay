package activitypub

import "github.com/go-fed/httpsig"

const (

	// ActorKeySize is the key size used to generate a new key
	ActorKeySize = 2048
	// ActorSummary is the description of the actor for the relay
	ActorSummary = "Feditools ActivityPub Relay - https://github.com/feditools/relay"

	// ContextActivityStreams contains the context document for activity streams
	ContextActivityStreams = "https://www.w3.org/ns/activitystreams"
)

var (
	validAlgs = []httpsig.Algorithm{
		httpsig.RSA_SHA512,
		httpsig.RSA_SHA256,
		httpsig.ED25519,
	}
)
