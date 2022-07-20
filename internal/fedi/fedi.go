package fedi

import (
	"context"
	"github.com/feditools/go-lib/fedihelper"
	"github.com/feditools/go-lib/fedihelper/mastodon"
	"github.com/feditools/relay/internal/config"
	"github.com/feditools/relay/internal/db"
	"github.com/feditools/relay/internal/kv"
	"github.com/feditools/relay/internal/models"
	"github.com/feditools/relay/internal/token"
	"github.com/spf13/viper"
	"net/url"
)

func New(d db.DB, t *fedihelper.Transport, k kv.KV, tok *token.Tokenizer) (*Module, error) {
	appName := viper.GetString(config.Keys.ApplicationName)
	appWebsite := viper.GetString(config.Keys.ApplicationWebsite)
	externalHostname := viper.GetString(config.Keys.ServerExternalHostname)

	// prep fedi helpers
	var fediHelpers []fedihelper.Helper
	mastoHelper, err := mastodon.New(k, t, appName, appWebsite, "https://"+externalHostname)
	if err != nil {
		return nil, err
	}
	fediHelpers = append(fediHelpers, mastoHelper)

	// prep fedi
	newModule := &Module{
		db:   d,
		tokz: tok,
	}

	fedi, err := fedihelper.New(k, t, appName, fediHelpers)
	if err != nil {
		return nil, err
	}

	fedi.SetCreateAccountHandler(newModule.CreateAccountHandler)
	fedi.SetGetAccountHandler(newModule.GetAccountHandler)
	fedi.SetNewAccountHandler(newModule.NewAccountHandler)

	newModule.helper = fedi

	return newModule, nil
}

type Module struct {
	db   db.DB
	tokz *token.Tokenizer

	helper *fedihelper.FediHelper
}

func (m *Module) FetchActor(ctx context.Context, actorIRI *url.URL) (*fedihelper.Actor, error) {
	return m.helper.FetchActor(ctx, actorIRI)
}

func (m *Module) FetchHostMeta(ctx context.Context, domain string) (*fedihelper.HostMeta, error) {
	return m.helper.FetchHostMeta(ctx, domain)
}

func (m *Module) FetchWebFinger(ctx context.Context, wfURI fedihelper.WebfingerURI, username, domain string) (*fedihelper.WebFinger, error) {
	return m.helper.FetchWebFinger(ctx, wfURI, username, domain)
}

func (m *Module) GetLoginURL(ctx context.Context, redirectURI *url.URL, instance *models.Instance) (*url.URL, error) {
	return m.helper.GetLoginURL(ctx, redirectURI, instance)
}

func (m *Module) GenInstanceFromDomain(ctx context.Context, domain string) (*models.Instance, error) {
	instance := new(models.Instance)

	err := m.helper.GenerateFediInstanceFromDomain(ctx, domain, instance)
	if err != nil {
		return nil, err
	}

	return instance, nil
}

func (m *Module) Helper(s fedihelper.SoftwareName) fedihelper.Helper {
	return m.helper.Helper(s)
}

func (m *Module) NewAccountFromUsername(ctx context.Context, username string, instance *models.Instance) (*models.Account, error) {
	account := new(models.Account)

	err := m.helper.GenerateFediAccountFromUsername(ctx, username, instance, account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (m *Module) NewInstanceFromDomain(ctx context.Context, domain string) (*models.Instance, error) {
	instance := new(models.Instance)

	err := m.helper.GenerateFediInstanceFromDomain(ctx, domain, instance)
	if err != nil {
		return nil, err
	}

	return instance, nil
}
