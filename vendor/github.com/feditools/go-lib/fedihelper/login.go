package fedihelper

import (
	"context"
	"net/url"

	"github.com/feditools/go-lib"
)

// GetLoginURL retrieves an oauth url for a federated instance.
func (f *FediHelper) GetLoginURL(ctx context.Context, act string) (*url.URL, error) {
	l := logger.WithField("func", "GetLoginURL")
	_, domain, err := lib.SplitAccount(act)
	if err != nil {
		l.Errorf("split account: %s", err.Error())

		return nil, err
	}

	// try to get instance from the database
	instance, found, err := f.GetInstanceHandler(ctx, domain)
	if err != nil {
		l.Errorf("db read: %s", err.Error())

		return nil, err
	}
	if found {
		u, err := f.loginURLForInstance(ctx, instance)
		if err != nil {
			l.Errorf("get login url: %s", err.Error())

			return nil, err
		}

		return u, nil
	}

	// get instance data from instance apis
	newInstance, err := f.NewInstanceHandler(ctx)
	if err != nil {
		fhErr := NewErrorf("new instance: %s", err.Error())
		l.Warn(fhErr.Error())

		return nil, fhErr
	}
	err = f.GenerateFediInstanceFromDomain(ctx, domain, newInstance)
	if err != nil {
		l.Errorf("get nodeinfo: %s", err.Error())

		return nil, err
	}
	err = f.CreateInstanceHandler(ctx, newInstance)
	if err != nil {
		l.Errorf("db create: %s", err.Error())

		return nil, err
	}

	u, err := f.loginURLForInstance(ctx, newInstance)
	if err != nil {
		l.Errorf("get login url: %s", err.Error())

		return nil, err
	}

	return u, nil
}

func (f *FediHelper) loginURLForInstance(ctx context.Context, instance Instance) (*url.URL, error) {
	l := logger.WithField("func", "loginURLForInstance")

	if _, ok := f.helpers[SoftwareName(instance.GetSoftware())]; !ok {
		return nil, NewErrorf("no helper for '%s'", instance.GetSoftware())
	}

	_, _, err := f.kv.GetInstanceOAuth(ctx, instance.GetID())
	if err != nil {
		if err.Error() != "nil" {
			fhErr := NewErrorf("kv get: %s", err.Error())
			l.Error(fhErr.Error())

			return nil, fhErr
		}

		var newClientID, newClientSecret string
		newClientID, newClientSecret, err = f.helpers[SoftwareMastodon].RegisterApp(ctx, instance)
		if err != nil {
			fhErr := NewErrorf("registering app: %s", err.Error())
			l.Error(fhErr.Error())

			return nil, fhErr
		}

		err = f.kv.SetInstanceOAuth(ctx, instance.GetID(), newClientID, newClientSecret)
		if err != nil {
			fhErr := NewErrorf("kv set: %s", err.Error())
			l.Error(fhErr.Error())

			return nil, fhErr
		}
	}

	return f.helpers[SoftwareMastodon].MakeLoginURI(ctx, instance)
}
