package mastodon

import (
	"context"
	"time"

	"github.com/feditools/go-lib/fedihelper"
)

// GetCurrentAccount retrieves the current federated account.
func (h *Helper) GetCurrentAccount(ctx context.Context, instance fedihelper.Instance, accessToken string) (fedihelper.Account, error) {
	l := logger.WithField("func", "GetCurrentAccount")

	// create mastodon client
	client, err := h.newClient(ctx, instance, accessToken)
	if err != nil {
		fhErr := fedihelper.NewErrorf("find actor url: %s", err.Error())
		l.Error(fhErr.Error())

		return nil, fhErr
	}

	// retrieve current account from
	retrievedAccount, err := client.GetAccountCurrentUser(ctx)
	if err != nil {
		fhErr := fedihelper.NewErrorf("getting current account: %s", err.Error())
		l.Error(fhErr.Error())

		return nil, fhErr
	}

	// check if account is locked
	if retrievedAccount.Locked {
		return nil, fedihelper.NewErrorf("account '@%s@%s' locked", retrievedAccount.Username, instance.GetDomain())
	}

	// check if account is a bot
	if retrievedAccount.Bot {
		return nil, fedihelper.NewErrorf("account '@%s@%s' is a bot", retrievedAccount.Username, instance.GetDomain())
	}

	// check if account has moved
	if retrievedAccount.Moved != nil {
		return nil, fedihelper.NewErrorf("account '@%s@%s' has moved to '@%s'", retrievedAccount.Username, instance.GetDomain(), retrievedAccount.Moved.Acct)
	}

	// try to retrieve federated account
	account, found, err := h.fedi.GetAccountHandler(ctx, instance, retrievedAccount.Username)
	if err != nil {
		fhErr := fedihelper.NewErrorf("get account: %s", err.Error())
		l.Error(fhErr.Error())

		return nil, fhErr
	}
	if found {
		return account, nil
	}

	// do webfinger
	webFinger, err := h.fedi.GetWellknownWebFinger(ctx, instance.GetServerHostname(), retrievedAccount.Username, instance.GetDomain())
	if err != nil {
		fhErr := fedihelper.NewErrorf("webfinger %s@%s: %s", retrievedAccount.Username, instance.GetDomain(), err.Error())
		l.Debug(fhErr.Error())

		return nil, fhErr
	}
	actorURI, err := fedihelper.FindActorURI(webFinger)
	if err != nil {
		fhErr := fedihelper.NewErrorf("finding actor uri %s@%s: %s", retrievedAccount.Username, instance.GetDomain(), err.Error())
		l.Debug(fhErr.Error())

		return nil, fhErr
	}
	if actorURI == nil {
		fhErr := fedihelper.NewErrorf("didn't find actor uri for %s@%s", retrievedAccount.Username, instance.GetDomain())
		l.Debug(fhErr.Error())

		return nil, fhErr
	}

	// create new federated account
	newFediAccount, err := h.fedi.NewAccountHandler(ctx)
	if err != nil {
		fhErr := fedihelper.NewErrorf("new account: %s", err.Error())
		l.Warn(fhErr.Error())

		return nil, fhErr
	}
	newFediAccount.SetActorURI(actorURI.String())
	newFediAccount.SetDisplayName(retrievedAccount.DisplayName)
	newFediAccount.SetInstance(instance)
	newFediAccount.SetLastFinger(time.Now())
	newFediAccount.SetUsername(retrievedAccount.Username)

	// write new federated account to database
	err = h.fedi.CreateAccountHandler(ctx, newFediAccount)
	if err != nil {
		fhErr := fedihelper.NewErrorf("db create: %s", err.Error())
		l.Error(fhErr.Error())

		return nil, fhErr
	}

	err = h.kv.SetAccessToken(ctx, newFediAccount.GetID(), accessToken)
	if err != nil {
		fhErr := fedihelper.NewErrorf("set access token: %s", err.Error())
		l.Error(fhErr.Error())

		return nil, fhErr
	}

	return newFediAccount, nil
}
