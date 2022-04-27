// Code generated by astool. DO NOT EDIT.

package propertypublickey

import vocab "github.com/go-fed/activity/streams/vocab"

var mgr privateManager

// privateManager abstracts the code-generated manager that provides access to
// concrete implementations.
type privateManager interface {
	// DeserializePublicKeyW3IDSecurityV1 returns the deserialization method
	// for the "W3IDSecurityV1PublicKey" non-functional property in the
	// vocabulary "W3IDSecurityV1"
	DeserializePublicKeyW3IDSecurityV1() func(map[string]interface{}, map[string]string) (vocab.W3IDSecurityV1PublicKey, error)
}

// SetManager sets the manager package-global variable. For internal use only, do
// not use as part of Application behavior. Must be called at golang init time.
func SetManager(m privateManager) {
	mgr = m
}
