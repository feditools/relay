// Code generated by astool. DO NOT EDIT.

package vocab

import "net/url"

// Specifies a Collection of actors who are working on the object, or responsible
// for it, or managing or administrating it, or having edit access to it. For
// example, for a Repository, it could be the people who have push/edit
// access, the "collaborators" of the repository.
//
//   {
//     "@context": [
//       "https://www.w3.org/ns/activitystreams",
//       "https://w3id.org/security/v1",
//       "https://forgefed.peers.community/ns"
//     ],
//     "followers": "https://dev.example/aviva/treesim/followers",
//     "id": "https://dev.example/aviva/treesim",
//     "inbox": "https://dev.example/aviva/treesim/inbox",
//     "name": "Tree Growth 3D Simulation",
//     "outbox": "https://dev.example/aviva/treesim/outbox",
//     "publicKey": {
//       "id": "https://dev.example/aviva/treesim#main-key",
//       "owner": "https://dev.example/aviva/treesim",
//       "publicKeyPem": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhki....."
//     },
//     "summary": "\u003cp\u003eTree growth 3D simulator for my nature
// exploration game\u003c/p\u003e",
//     "team": "https://dev.example/aviva/treesim/team",
//     "type": "Repository"
//   }
//
//   {
//     "@context": "https://www.w3.org/ns/activitystreams",
//     "id": "https://dev.example/aviva/treesim/team",
//     "items": [
//       "https://dev.example/aviva",
//       "https://dev.example/luke",
//       "https://code.community/users/lorax"
//     ],
//     "totalItems": 3,
//     "type": "Collection"
//   }
type ForgeFedTeamProperty interface {
	// Clear ensures no value of this property is set. Calling HasAny or any
	// of the 'Is' methods afterwards will return false.
	Clear()
	// GetActivityStreamsCollection returns the value of this property. When
	// IsActivityStreamsCollection returns false,
	// GetActivityStreamsCollection will return an arbitrary value.
	GetActivityStreamsCollection() ActivityStreamsCollection
	// GetActivityStreamsCollectionPage returns the value of this property.
	// When IsActivityStreamsCollectionPage returns false,
	// GetActivityStreamsCollectionPage will return an arbitrary value.
	GetActivityStreamsCollectionPage() ActivityStreamsCollectionPage
	// GetActivityStreamsOrderedCollection returns the value of this property.
	// When IsActivityStreamsOrderedCollection returns false,
	// GetActivityStreamsOrderedCollection will return an arbitrary value.
	GetActivityStreamsOrderedCollection() ActivityStreamsOrderedCollection
	// GetActivityStreamsOrderedCollectionPage returns the value of this
	// property. When IsActivityStreamsOrderedCollectionPage returns
	// false, GetActivityStreamsOrderedCollectionPage will return an
	// arbitrary value.
	GetActivityStreamsOrderedCollectionPage() ActivityStreamsOrderedCollectionPage
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return an arbitrary value.
	GetIRI() *url.URL
	// GetType returns the value in this property as a Type. Returns nil if
	// the value is not an ActivityStreams type, such as an IRI or another
	// value.
	GetType() Type
	// HasAny returns true if any of the different values is set.
	HasAny() bool
	// IsActivityStreamsCollection returns true if this property has a type of
	// "Collection". When true, use the GetActivityStreamsCollection and
	// SetActivityStreamsCollection methods to access and set this
	// property.
	IsActivityStreamsCollection() bool
	// IsActivityStreamsCollectionPage returns true if this property has a
	// type of "CollectionPage". When true, use the
	// GetActivityStreamsCollectionPage and
	// SetActivityStreamsCollectionPage methods to access and set this
	// property.
	IsActivityStreamsCollectionPage() bool
	// IsActivityStreamsOrderedCollection returns true if this property has a
	// type of "OrderedCollection". When true, use the
	// GetActivityStreamsOrderedCollection and
	// SetActivityStreamsOrderedCollection methods to access and set this
	// property.
	IsActivityStreamsOrderedCollection() bool
	// IsActivityStreamsOrderedCollectionPage returns true if this property
	// has a type of "OrderedCollectionPage". When true, use the
	// GetActivityStreamsOrderedCollectionPage and
	// SetActivityStreamsOrderedCollectionPage methods to access and set
	// this property.
	IsActivityStreamsOrderedCollectionPage() bool
	// IsIRI returns true if this property is an IRI. When true, use GetIRI
	// and SetIRI to access and set this property
	IsIRI() bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// KindIndex computes an arbitrary value for indexing this kind of value.
	// This is a leaky API detail only for folks looking to replace the
	// go-fed implementation. Applications should not use this method.
	KindIndex() int
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o ForgeFedTeamProperty) bool
	// Name returns the name of this property: "team".
	Name() string
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// SetActivityStreamsCollection sets the value of this property. Calling
	// IsActivityStreamsCollection afterwards returns true.
	SetActivityStreamsCollection(v ActivityStreamsCollection)
	// SetActivityStreamsCollectionPage sets the value of this property.
	// Calling IsActivityStreamsCollectionPage afterwards returns true.
	SetActivityStreamsCollectionPage(v ActivityStreamsCollectionPage)
	// SetActivityStreamsOrderedCollection sets the value of this property.
	// Calling IsActivityStreamsOrderedCollection afterwards returns true.
	SetActivityStreamsOrderedCollection(v ActivityStreamsOrderedCollection)
	// SetActivityStreamsOrderedCollectionPage sets the value of this
	// property. Calling IsActivityStreamsOrderedCollectionPage afterwards
	// returns true.
	SetActivityStreamsOrderedCollectionPage(v ActivityStreamsOrderedCollectionPage)
	// SetIRI sets the value of this property. Calling IsIRI afterwards
	// returns true.
	SetIRI(v *url.URL)
	// SetType attempts to set the property for the arbitrary type. Returns an
	// error if it is not a valid type to set on this property.
	SetType(t Type) error
}
