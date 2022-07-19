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

func (m *Module) CreateInstanceHandler(ctx context.Context, instanceI fedihelper.Instance) (err error) {
	instance, ok := instanceI.(*models.Instance)
	if !ok {
		return ErrCantCast
	}

	return m.db.CreateInstance(ctx, instance)
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

func (m *Module) GetInstanceHandler(ctx context.Context, domain string) (fedihelper.Instance, bool, error) {
	instance, err := m.db.ReadInstanceByDomain(ctx, domain)
	if err != nil {
		if errors.Is(err, db.ErrNoEntries) {
			return nil, false, nil
		}

		return nil, false, err
	}

	return instance, true, nil
}

func (m *Module) GetTokenHandler(_ context.Context, o interface{}) (token string) {
	return m.tokz.GetToken(o)
}

func (*Module) NewAccountHandler(_ context.Context) (account fedihelper.Account, err error) {
	return &models.Account{}, nil
}

func (*Module) NewInstanceHandler(_ context.Context) (instance fedihelper.Instance, err error) {
	return &models.Instance{}, nil
}

func (m *Module) UpdateInstanceHandler(ctx context.Context, instanceI fedihelper.Instance) (err error) {
	instance, ok := instanceI.(*models.Instance)
	if !ok {
		return ErrCantCast
	}

	return m.db.UpdateInstance(ctx, instance)
}
