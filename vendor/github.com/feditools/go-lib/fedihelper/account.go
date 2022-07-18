package fedihelper

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/feditools/go-lib/fedihelper/models"
	httplib "github.com/feditools/go-lib/http"
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
	hostMeta, err := f.GetHostMeta(ctx, instance.GetDomain())
	if err != nil {
		l.Errorf("get host meta: %s", err.Error())

		return err
	}
	hostMetaFTemplate, err := f.WebfingerURIFromHostMeta(hostMeta)
	if err != nil {
		l.Errorf("get webfinger uri: %s", err.Error())

		return err
	}

	// get actor uri
	webfinger, err := f.webFinger(ctx, hostMetaFTemplate, username, instance.GetDomain())
	if err != nil {
		fhErr := NewErrorf("get wellknown webfinger: %s", err.Error())
		l.Error(fhErr.Error())

		return fhErr
	}
	actorURI, err := FindActorURI(webfinger)
	if err != nil {
		fhErr := NewErrorf("find actor url: %s", err.Error())
		l.Error(fhErr.Error())

		return fhErr
	}
	if actorURI == nil {
		return NewError("missing actor uri")
	}

	// retrieve actor
	v, err, _ := f.requestGroup.Do(actorURI.String(), func() (interface{}, error) {
		// do request
		req, err := f.http.NewRequest(ctx, http.MethodGet, actorURI.String(), nil)
		req.Header.Add(httplib.HeaderAccept, string(httplib.MimeAppJSON))
		if err != nil {
			l.Errorf("new http request: %s", err.Error())

			return nil, err
		}
		resp, err := f.http.Do(req)
		if err != nil {
			l.Errorf("http do: %s", err.Error())

			return nil, err
		}

		actorinfo := new(models.Actor)
		defer resp.Body.Close()
		err = json.NewDecoder(resp.Body).Decode(actorinfo)
		if err != nil {
			l.Errorf("decode json: %s", err.Error())

			return nil, err
		}

		return actorinfo, nil
	})

	if err != nil {
		fhErr := NewErrorf("singleflight: %s", err.Error())
		l.Error(fhErr.Error())

		return fhErr
	}

	actorinfo, ok := v.(*models.Actor)
	if !ok {
		return NewError("invalid response type from single flight")
	}

	if actorinfo.Type != "Person" {
		return NewError("actor is not a person")
	}

	account.SetActorURI(actorURI.String())
	account.SetUsername(actorinfo.PreferredUsername)
	account.SetInstance(instance)
	account.SetDisplayName(actorinfo.Name)
	account.SetLastFinger(time.Now())

	return nil
}
