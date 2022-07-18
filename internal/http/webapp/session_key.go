package webapp

// SessionKey is a key used for storing data in a web session.
type SessionKey int

const (
	// SessionKeyAccountID contains the id of the currently logged-in user.
	SessionKeyAccountID SessionKey = iota
)
