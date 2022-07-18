package redis

import (
	"context"
	"encoding/json"
	"github.com/feditools/relay/internal/kv"
	"github.com/feditools/relay/internal/util"
)

type instanceOAuth struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (c *Client) DeleteInstanceOAuth(ctx context.Context, instanceID int64) error {
	_, err := c.redis.Del(ctx, kv.KeyInstanceOAuth(instanceID)).Result()
	if err != nil {
		return c.ProcessError(err)
	}

	return nil
}

func (c *Client) GetInstanceOAuth(ctx context.Context, instanceID int64) (string, string, error) {
	resp, err := c.redis.Get(ctx, kv.KeyInstanceOAuth(instanceID)).Bytes()
	if err != nil {
		return "", "", c.ProcessError(err)
	}

	jsondata, err := util.Decrypt(resp)
	if err != nil {
		return "", "", kv.NewEncryptionError(err.Error())
	}

	var io instanceOAuth
	err = json.Unmarshal(jsondata, &io)
	if err != nil {
		return "", "", c.ProcessError(err)
	}

	return io.ClientID, io.ClientSecret, nil
}

func (c *Client) SetInstanceOAuth(ctx context.Context, instanceID int64, clientID string, clientSecret string) error {
	io := instanceOAuth{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}

	jsonData, err := json.Marshal(io)
	if err != nil {
		return c.ProcessError(err)
	}

	data, err := util.Encrypt(jsonData)
	if err != nil {
		return kv.NewEncryptionError(err.Error())
	}

	_, err = c.redis.Set(ctx, kv.KeyInstanceOAuth(instanceID), data, 0).Result()
	if err != nil {
		return c.ProcessError(err)
	}

	return nil
}
