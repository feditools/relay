package redis

import (
	"errors"
	"github.com/feditools/relay/internal/kv"
	"github.com/go-redis/redis/v8"
)

// ProcessError replaces any known values with our own db.Error types.
func (*Client) ProcessError(err error) kv.Error {
	switch {
	case err == nil:
		return nil
	case errors.Is(err, redis.Nil):
		return kv.ErrNil
	default:
		return err
	}
}
