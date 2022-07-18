package kv

import (
	"github.com/feditools/go-lib/fedihelper"
)

// KV represents a key value store.
type KV interface {
	fedihelper.KV
}
