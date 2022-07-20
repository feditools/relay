package fedihelper

import (
	"context"
	"net/url"
)

// GetLoginURL retrieves an oauth url for a federated instance.
func (f *FediHelper) GetLoginURL(ctx context.Context, redirectURI *url.URL, instance Instance) (*url.URL, error) {
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
		newClientID, newClientSecret, err = f.helpers[SoftwareMastodon].RegisterApp(ctx, redirectURI, instance)
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

	return f.helpers[SoftwareMastodon].MakeLoginURI(ctx, redirectURI, instance)
}

/*func (f *FediHelper) loginURLForInstance(ctx context.Context, redirectURI *url.URL, instance Instance) (*url.URL, error) {
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
		newClientID, newClientSecret, err = f.helpers[SoftwareMastodon].RegisterApp(ctx, redirectURI, instance)
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

	return f.helpers[SoftwareMastodon].MakeLoginURI(ctx, redirectURI, instance)
}*/
