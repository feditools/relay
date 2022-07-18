package redis

import (
	"context"
	"github.com/feditools/relay/internal/kv"
	"github.com/feditools/relay/internal/util"
)

func (c *Client) DeleteAccessToken(ctx context.Context, accountID int64) error {
	_, err := c.redis.Del(ctx, kv.KeyAccountAccessToken(accountID)).Result()
	if err != nil {
		return c.ProcessError(err)
	}

	return nil
}

func (c *Client) GetAccessToken(ctx context.Context, accountID int64) (string, error) {
	resp, err := c.redis.Get(ctx, kv.KeyAccountAccessToken(accountID)).Bytes()
	if err != nil {
		return "", c.ProcessError(err)
	}

	data, err := util.Decrypt(resp)
	if err != nil {
		return "", kv.NewEncryptionError(err.Error())
	}

	return string(data), nil
}

func (c *Client) SetAccessToken(ctx context.Context, accountID int64, accessToken string) error {
	data, err := util.Encrypt([]byte(accessToken))
	if err != nil {
		return kv.NewEncryptionError(err.Error())
	}

	_, err = c.redis.Set(ctx, kv.KeyAccountAccessToken(accountID), data, 0).Result()
	if err != nil {
		return c.ProcessError(err)
	}

	return nil
}
