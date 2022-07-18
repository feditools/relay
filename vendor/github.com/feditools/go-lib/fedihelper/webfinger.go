package fedihelper

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/feditools/go-lib/fedihelper/models"
)

// webFinger retrieves web finger resource from a federated instance.
func (f *FediHelper) webFinger(ctx context.Context, fTemplate, username, domain string) (*models.WebFinger, error) {
	l := logger.WithField("func", "webFinger")
	webfingerURI := fmt.Sprintf(fTemplate, username, domain)
	v, err, _ := f.requestGroup.Do(webfingerURI, func() (interface{}, error) {
		l.Tracef("webfingering %s", webfingerURI)

		// do request
		resp, err := f.http.Get(ctx, webfingerURI)
		if err != nil {
			l.Errorf("http get: %s", err.Error())

			return nil, err
		}

		webfinger := new(models.WebFinger)
		defer resp.Body.Close()
		err = json.NewDecoder(resp.Body).Decode(webfinger)
		if err != nil {
			l.Errorf("decode json: %s", err.Error())

			return nil, err
		}

		return webfinger, nil
	})

	if err != nil {
		l.Errorf("singleflight: %s", err.Error())

		return nil, err
	}

	webfinger, ok := v.(*models.WebFinger)
	if !ok {
		return nil, NewError("invalid response type from single flight")
	}

	return webfinger, nil
}
