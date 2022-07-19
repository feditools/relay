package token

// Kind represents the kind of model to encode a token for.
type Kind int64

// This order can not change else all external urls with tokens will become invalid.
const (
	// KindInstance is a token that represents a federated social instance.
	KindInstance Kind = 1 + iota
	// KindAccount is a token that represents a federated social account.
	KindAccount
	// KindBlock is a token that represents a blocked federated social instance.
	KindBlock
)

func (k Kind) String() string {
	switch k {
	case KindInstance:
		return "Instance"
	case KindAccount:
		return "Account"
	case KindBlock:
		return "Block"
	default:
		return "unknown"
	}
}
