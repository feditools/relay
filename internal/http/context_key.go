package http

// ContextKey is a key used in http request contexts
type ContextKey int

const (
	// ContextKeyKeyVerifier is a http signature verifier
	ContextKeyKeyVerifier ContextKey = iota
	// ContextKeyHTTPSignature is a http signature
	ContextKeyHTTPSignature
)
