package redis

import (
	"context"
	"github.com/feditools/relay/internal/config"
	"github.com/feditools/relay/internal/kv"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

// New creates a new redis client.
func New(ctx context.Context) (*Client, error) {
	l := logger.WithField("func", "New")

	r := redis.NewClient(&redis.Options{
		Addr:     viper.GetString(config.Keys.RedisAddress),
		Password: viper.GetString(config.Keys.RedisPassword),
		DB:       viper.GetInt(config.Keys.RedisDB),
	})

	c := Client{
		redis: r,
	}

	resp := c.redis.Ping(ctx)
	l.Debugf("%s", resp.String())

	return &c, nil
}

// Client represents a redis client.
type Client struct {
	redis *redis.Client
}

var _ kv.KV = (*Client)(nil)

// Close closes the redis pool.
func (c *Client) Close(_ context.Context) kv.Error {
	return c.redis.Close()
}

// RedisClient returns the redis client.
func (c *Client) RedisClient() *redis.Client {
	return c.redis
}
