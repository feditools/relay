// Code generated by astool. DO NOT EDIT.

package propertydependedby

import vocab "github.com/go-fed/activity/streams/vocab"

var mgr privateManager

// privateManager abstracts the code-generated manager that provides access to
// concrete implementations.
type privateManager interface {
	// DeserializeTicketForgeFed returns the deserialization method for the
	// "ForgeFedTicket" non-functional property in the vocabulary
	// "ForgeFed"
	DeserializeTicketForgeFed() func(map[string]interface{}, map[string]string) (vocab.ForgeFedTicket, error)
}

// SetManager sets the manager package-global variable. For internal use only, do
// not use as part of Application behavior. Must be called at golang init time.
func SetManager(m privateManager) {
	mgr = m
}
