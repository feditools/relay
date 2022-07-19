package fedihelper

import (
	"context"
	"time"
)

type Account interface {
	GetActorURI() (actorURI string)
	GetDisplayName() (displayName string)
	GetID() (id int64)
	GetInstance() (instance Instance)
	GetLastFinger() (lastFinger time.Time)
	GetUsername() (username string)

	SetActorURI(actorURI string)
	SetDisplayName(displayName string)
	SetInstance(instance Instance)
	SetLastFinger(lastFinger time.Time)
	SetUsername(username string)
}

// GenerateFediAccountFromUsername creates an Account object by querying the apis of the federated instance.
func (f *FediHelper) GenerateFediAccountFromUsername(ctx context.Context, username string, instance Instance, account Account) error {
	l := logger.WithField("func", "GenerateFediAccountFromUsername")

	// get host meta
	hostMeta, err := f.FetchHostMeta(ctx, instance.GetDomain())
	if err != nil {
		l.Errorf("get host meta: %s", err.Error())

		return err
	}
	webfingerURI := hostMeta.WebfingerURI()
	if webfingerURI == "" {
		l.Errorf("host meta missing web finger url")

		return NewError("host meta missing web finger url")
	}

	// get actor uri
	webfinger, err := f.FetchWebFinger(ctx, webfingerURI, username, instance.GetDomain())
	if err != nil {
		fhErr := NewErrorf("get wellknown webfinger: %s", err.Error())
		l.Error(fhErr.Error())

		return fhErr
	}
	actorURI, err := webfinger.ActorURI()
	if err != nil {
		fhErr := NewErrorf("find actor url: %s", err.Error())
		l.Error(fhErr.Error())

		return fhErr
	}
	if actorURI == nil {
		return NewError("missing actor uri")
	}

	actor, err := f.FetchActor(ctx, actorURI)
	if err != nil {
		l.Errorf("decode json: %s", err.Error())

		return err
	}

	account.SetActorURI(actorURI.String())
	account.SetUsername(actor.PreferredUsername)
	account.SetInstance(instance)
	account.SetDisplayName(actor.Name)
	account.SetLastFinger(time.Now())

	return nil
}
