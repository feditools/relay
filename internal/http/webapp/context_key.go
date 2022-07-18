package webapp

// ContextKey is a key used in http request contexts.
type ContextKey int

const (
	// ContextKeySession is the persistent session.
	ContextKeySession ContextKey = iota
	// ContextKeyLocalizer is the language localizer.
	ContextKeyLocalizer
	// ContextKeyLanguage is the language.
	ContextKeyLanguage
	// ContextKeyAccount is the logged in user's account.
	ContextKeyAccount
	// ContextKeyOauthNonce is the oauth nonce.
	ContextKeyOauthNonce
)
