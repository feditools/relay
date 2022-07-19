package redis

import (
	"context"
	"github.com/feditools/relay/internal/kv"
	"time"
)

// DeleteActor deletes fedi actor from redis.
func (c *Client) DeleteActor(ctx context.Context, actorURI string) error {
	_, err := c.redis.Del(ctx, kv.KeyFediActor(actorURI)).Result()
	if err != nil {
		return c.ProcessError(err)
	}

	return nil
}

// GetActor retrieves fedi actor from redis.
func (c *Client) GetActor(ctx context.Context, actorURI string) ([]byte, error) {
	resp, err := c.redis.Get(ctx, kv.KeyFediActor(actorURI)).Bytes()
	if err != nil {
		return nil, c.ProcessError(err)
	}

	return resp, nil
}

// SetActor adds fedi actor to redis.
func (c *Client) SetActor(ctx context.Context, actorURI string, actor []byte, expire time.Duration) error {
	_, err := c.redis.SetEX(ctx, kv.KeyFediActor(actorURI), actor, expire).Result()
	if err != nil {
		return c.ProcessError(err)
	}

	return nil
}
