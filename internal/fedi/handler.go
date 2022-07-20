package fedi

import (
	"context"
	"errors"
	"github.com/feditools/go-lib/fedihelper"
	"github.com/feditools/relay/internal/db"
	"github.com/feditools/relay/internal/models"
)

func (m *Module) CreateAccountHandler(ctx context.Context, accountI fedihelper.Account) (err error) {
	account, ok := accountI.(*models.Account)
	if !ok {
		return ErrCantCast
	}

	return m.db.CreateAccount(ctx, account)
}

func (m *Module) GetAccountHandler(ctx context.Context, instanceI fedihelper.Instance, username string) (fedihelper.Account, bool, error) {
	instance, ok := instanceI.(*models.Instance)
	if !ok {
		return nil, false, ErrCantCast
	}

	account, err := m.db.ReadAccountByUsername(ctx, instance.ID, username)
	if err != nil {
		if errors.Is(err, db.ErrNoEntries) {
			return nil, false, nil
		}

		return nil, false, err
	}

	return account, true, nil
}

func (*Module) NewAccountHandler(_ context.Context) (account fedihelper.Account, err error) {
	return &models.Account{}, nil
}
