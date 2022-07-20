package fedihelper

import "context"

type (
	CreateAccountHandler func(ctx context.Context, account Account) (err error)
	GetAccountHandler    func(ctx context.Context, instance Instance, username string) (account Account, found bool, err error)
	NewAccountHandler    func(ctx context.Context) (account Account, err error)
)
