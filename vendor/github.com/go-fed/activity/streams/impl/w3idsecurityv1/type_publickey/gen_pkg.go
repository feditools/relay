// Code generated by astool. DO NOT EDIT.

package typepublickey

import vocab "github.com/go-fed/activity/streams/vocab"

var mgr privateManager

var typePropertyConstructor func() vocab.JSONLDTypeProperty

// privateManager abstracts the code-generated manager that provides access to
// concrete implementations.
type privateManager interface {
	// DeserializeIdPropertyJSONLD returns the deserialization method for the
	// "JSONLDIdProperty" non-functional property in the vocabulary
	// "JSONLD"
	DeserializeIdPropertyJSONLD() func(map[string]interface{}, map[string]string) (vocab.JSONLDIdProperty, error)
	// DeserializeOwnerPropertyW3IDSecurityV1 returns the deserialization
	// method for the "W3IDSecurityV1OwnerProperty" non-functional
	// property in the vocabulary "W3IDSecurityV1"
	DeserializeOwnerPropertyW3IDSecurityV1() func(map[string]interface{}, map[string]string) (vocab.W3IDSecurityV1OwnerProperty, error)
	// DeserializePublicKeyPemPropertyW3IDSecurityV1 returns the
	// deserialization method for the "W3IDSecurityV1PublicKeyPemProperty"
	// non-functional property in the vocabulary "W3IDSecurityV1"
	DeserializePublicKeyPemPropertyW3IDSecurityV1() func(map[string]interface{}, map[string]string) (vocab.W3IDSecurityV1PublicKeyPemProperty, error)
}

// jsonldContexter is a private interface to determine the JSON-LD contexts and
// aliases needed for functional and non-functional properties. It is a helper
// interface for this implementation.
type jsonldContexter interface {
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
}

// SetManager sets the manager package-global variable. For internal use only, do
// not use as part of Application behavior. Must be called at golang init time.
func SetManager(m privateManager) {
	mgr = m
}

// SetTypePropertyConstructor sets the "type" property's constructor in the
// package-global variable. For internal use only, do not use as part of
// Application behavior. Must be called at golang init time. Permits
// ActivityStreams types to correctly set their "type" property at
// construction time, so users don't have to remember to do so each time. It
// is dependency injected so other go-fed compatible implementations could
// inject their own type.
func SetTypePropertyConstructor(f func() vocab.JSONLDTypeProperty) {
	typePropertyConstructor = f
}
