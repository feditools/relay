package fedihelper

import (
	"context"
	"time"
)

type KV interface {
	// access token

	DeleteAccessToken(ctx context.Context, accountID int64) (err error)
	GetAccessToken(ctx context.Context, accountID int64) (accessToken string, err error)
	SetAccessToken(ctx context.Context, accountID int64, accessToken string) (err error)

	// access token

	DeleteActor(ctx context.Context, actorURI string) (err error)
	GetActor(ctx context.Context, actorURI string) (actor []byte, err error)
	SetActor(ctx context.Context, actorURI string, actor []byte, expire time.Duration) (err error)

	// federated instance node info

	DeleteHostMeta(ctx context.Context, domain string) (err error)
	GetHostMeta(ctx context.Context, domain string) (hostmeta []byte, err error)
	SetHostMeta(ctx context.Context, domain string, hostmeta []byte, expire time.Duration) (err error)

	// instance oauth

	DeleteInstanceOAuth(ctx context.Context, instanceID int64) (err error)
	GetInstanceOAuth(ctx context.Context, instanceID int64) (clientID string, clientSecret string, err error)
	SetInstanceOAuth(ctx context.Context, instanceID int64, clientID string, clientSecret string) (err error)

	// federated instance node info

	DeleteFediNodeInfo(ctx context.Context, domain string) (err error)
	GetFediNodeInfo(ctx context.Context, domain string) (nodeinfo []byte, err error)
	SetFediNodeInfo(ctx context.Context, domain string, nodeinfo []byte, expire time.Duration) (err error)
}
