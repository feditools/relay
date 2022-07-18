// Code generated by astool. DO NOT EDIT.

package vocab

import "net/url"

// ActivityStreamsContentPropertyIterator represents a single value for the
// "content" property.
type ActivityStreamsContentPropertyIterator interface {
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return an arbitrary value.
	GetIRI() *url.URL
	// GetLanguage returns the value for the specified BCP47 language code, or
	// an empty string if it is either not a language map or no value is
	// present.
	GetLanguage(bcp47 string) string
	// GetRDFLangString returns the value of this property. When
	// IsRDFLangString returns false, GetRDFLangString will return an
	// arbitrary value.
	GetRDFLangString() map[string]string
	// GetXMLSchemaString returns the value of this property. When
	// IsXMLSchemaString returns false, GetXMLSchemaString will return an
	// arbitrary value.
	GetXMLSchemaString() string
	// HasAny returns true if any of the values are set, except for the
	// natural language map. When true, the specific has, getter, and
	// setter methods may be used to determine what kind of value there is
	// to access and set this property. To determine if the property was
	// set as a natural language map, use the IsRDFLangString method
	// instead.
	HasAny() bool
	// HasLanguage returns true if the natural language map has an entry for
	// the specified BCP47 language code.
	HasLanguage(bcp47 string) bool
	// IsIRI returns true if this property is an IRI. When true, use GetIRI
	// and SetIRI to access and set this property
	IsIRI() bool
	// IsRDFLangString returns true if this property has a type of
	// "langString". When true, use the GetRDFLangString and
	// SetRDFLangString methods to access and set this property.. To
	// determine if the property was set as a natural language map, use
	// the IsRDFLangString method instead.
	IsRDFLangString() bool
	// IsXMLSchemaString returns true if this property has a type of "string".
	// When true, use the GetXMLSchemaString and SetXMLSchemaString
	// methods to access and set this property.. To determine if the
	// property was set as a natural language map, use the IsRDFLangString
	// method instead.
	IsXMLSchemaString() bool
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
	LessThan(o ActivityStreamsContentPropertyIterator) bool
	// Name returns the name of this property: "ActivityStreamsContent".
	Name() string
	// Next returns the next iterator, or nil if there is no next iterator.
	Next() ActivityStreamsContentPropertyIterator
	// Prev returns the previous iterator, or nil if there is no previous
	// iterator.
	Prev() ActivityStreamsContentPropertyIterator
	// SetIRI sets the value of this property. Calling IsIRI afterwards
	// returns true.
	SetIRI(v *url.URL)
	// SetLanguage sets the value for the specified BCP47 language code.
	SetLanguage(bcp47, value string)
	// SetRDFLangString sets the value of this property and clears the natural
	// language map. Calling IsRDFLangString afterwards will return true.
	// Calling IsRDFLangString afterwards returns false.
	SetRDFLangString(v map[string]string)
	// SetXMLSchemaString sets the value of this property and clears the
	// natural language map. Calling IsXMLSchemaString afterwards will
	// return true. Calling IsRDFLangString afterwards returns false.
	SetXMLSchemaString(v string)
}

// The content or textual representation of the Object encoded as a JSON string.
// By default, the value of content is HTML. The mediaType property can be
// used in the object to indicate a different content type. The content MAY be
// expressed using multiple language-tagged values.
//
// Example 114 (https://www.w3.org/TR/activitystreams-vocabulary/#ex130-jsonld):
//   {
//     "content": "A \u003cem\u003esimple\u003c/em\u003e note",
//     "summary": "A simple note",
//     "type": "Note"
//   }
//
// Example 115 (https://www.w3.org/TR/activitystreams-vocabulary/#ex131-jsonld):
//   {
//     "contentMap": {
//       "en": "A \u003cem\u003esimple\u003c/em\u003e note",
//       "es": "Una nota \u003cem\u003esencilla\u003c/em\u003e",
//       "zh-hans": "一段\u003cem\u003e简单的\u003c/em\u003e笔记"
//     },
//     "summary": "A simple note",
//     "type": "Note"
//   }
//
// Example 116 (https://www.w3.org/TR/activitystreams-vocabulary/#ex130b-jsonld):
//   {
//     "content": "## A simple note\nA simple markdown `note`",
//     "mediaType": "text/markdown",
//     "summary": "A simple note",
//     "type": "Note"
//   }
type ActivityStreamsContentProperty interface {
	// AppendIRI appends an IRI value to the back of a list of the property
	// "content"
	AppendIRI(v *url.URL)
	// AppendRDFLangString appends a langString value to the back of a list of
	// the property "content". Invalidates iterators that are traversing
	// using Prev.
	AppendRDFLangString(v map[string]string)
	// AppendXMLSchemaString appends a string value to the back of a list of
	// the property "content". Invalidates iterators that are traversing
	// using Prev.
	AppendXMLSchemaString(v string)
	// At returns the property value for the specified index. Panics if the
	// index is out of bounds.
	At(index int) ActivityStreamsContentPropertyIterator
	// Begin returns the first iterator, or nil if empty. Can be used with the
	// iterator's Next method and this property's End method to iterate
	// from front to back through all values.
	Begin() ActivityStreamsContentPropertyIterator
	// Empty returns returns true if there are no elements.
	Empty() bool
	// End returns beyond-the-last iterator, which is nil. Can be used with
	// the iterator's Next method and this property's Begin method to
	// iterate from front to back through all values.
	End() ActivityStreamsContentPropertyIterator
	// Insert inserts an IRI value at the specified index for a property
	// "content". Existing elements at that index and higher are shifted
	// back once. Invalidates all iterators.
	InsertIRI(idx int, v *url.URL)
	// InsertRDFLangString inserts a langString value at the specified index
	// for a property "content". Existing elements at that index and
	// higher are shifted back once. Invalidates all iterators.
	InsertRDFLangString(idx int, v map[string]string)
	// InsertXMLSchemaString inserts a string value at the specified index for
	// a property "content". Existing elements at that index and higher
	// are shifted back once. Invalidates all iterators.
	InsertXMLSchemaString(idx int, v string)
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// KindIndex computes an arbitrary value for indexing this kind of value.
	// This is a leaky API method specifically needed only for alternate
	// implementations for go-fed. Applications should not use this
	// method. Panics if the index is out of bounds.
	KindIndex(idx int) int
	// Len returns the number of values that exist for the "content" property.
	Len() (length int)
	// Less computes whether another property is less than this one. Mixing
	// types results in a consistent but arbitrary ordering
	Less(i, j int) bool
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o ActivityStreamsContentProperty) bool
	// Name returns the name of this property ("content") with any alias.
	Name() string
	// PrependIRI prepends an IRI value to the front of a list of the property
	// "content".
	PrependIRI(v *url.URL)
	// PrependRDFLangString prepends a langString value to the front of a list
	// of the property "content". Invalidates all iterators.
	PrependRDFLangString(v map[string]string)
	// PrependXMLSchemaString prepends a string value to the front of a list
	// of the property "content". Invalidates all iterators.
	PrependXMLSchemaString(v string)
	// Remove deletes an element at the specified index from a list of the
	// property "content", regardless of its type. Panics if the index is
	// out of bounds. Invalidates all iterators.
	Remove(idx int)
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// SetIRI sets an IRI value to be at the specified index for the property
	// "content". Panics if the index is out of bounds.
	SetIRI(idx int, v *url.URL)
	// SetRDFLangString sets a langString value to be at the specified index
	// for the property "content". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetRDFLangString(idx int, v map[string]string)
	// SetXMLSchemaString sets a string value to be at the specified index for
	// the property "content". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetXMLSchemaString(idx int, v string)
	// Swap swaps the location of values at two indices for the "content"
	// property.
	Swap(i, j int)
}
