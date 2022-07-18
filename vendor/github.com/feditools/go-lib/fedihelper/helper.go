package fedihelper

import (
	"context"
	"net/url"
)

// Helper interacts with a federated social instance.
type Helper interface {
	GetAccessToken(ctx context.Context, instance Instance, code string) (accessToken string, err error)
	GetCurrentAccount(ctx context.Context, instance Instance, accessToken string) (user Account, err error)
	GetSoftware() Software
	RegisterApp(ctx context.Context, instance Instance) (clientID string, clientSecret string, err error)
	SetFedi(f *FediHelper)
	MakeLoginURI(ctx context.Context, instance Instance) (loginURI *url.URL, err error)
}

// Helper returns a helper for a given software package.
func (f *FediHelper) Helper(s Software) Helper {
	return f.helpers[s]
}
