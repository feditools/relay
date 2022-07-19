package fedihelper

import (
	"context"
	"encoding/json"
	"fmt"
	libhttp "github.com/feditools/go-lib/http"
	"net/url"
)

// WebFinger is a web finger response.
type WebFinger struct {
	Aliases []string `json:"aliases,omitempty"`
	Links   []Link   `json:"links,omitempty"`
	Subject string   `json:"subject,omitempty"`
}

// ActorURI an actor uri.
func (w *WebFinger) ActorURI() (*url.URL, error) {
	var actorURIstr string
	for _, link := range w.Links {
		if link.Rel == "self" || link.Type == libhttp.MimeAppActivityJSON.String() {
			actorURIstr = link.Href

			break
		}
	}
	if actorURIstr == "" {
		return nil, NewError("missing actor uri")
	}

	actorURI, err := url.Parse(actorURIstr)
	if err != nil {
		return nil, NewErrorf("invalid actor uri: %s", err.Error())
	}

	return actorURI, err
}

// FetchWebFinger retrieves web finger resource from a federated instance.
func (f *FediHelper) FetchWebFinger(ctx context.Context, wfURI WebfingerURI, username, domain string) (*WebFinger, error) {
	l := logger.WithField("func", "FetchWebFinger")
	webFingerString := fmt.Sprintf(wfURI.FTemplate(), username, domain)
	webFingerURI, err := url.Parse(webFingerString)
	if err != nil {
		l.Errorf("parsing url '%s': %s", webFingerString, err.Error())

		return nil, err
	}

	v, err, _ := f.requestGroup.Do(fmt.Sprintf("webfinger-%s", webFingerURI.String()), func() (interface{}, error) {
		l.Tracef("webfingering %s", webFingerURI.String())

		// do request
		body, err := f.http.InstanceGet(ctx, webFingerURI)
		if err != nil {
			l.Errorf("http get: %s", err.Error())

			return nil, err
		}

		webfinger := new(WebFinger)
		err = json.Unmarshal(body, webfinger)
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

	webfinger, ok := v.(*WebFinger)
	if !ok {
		return nil, NewError("invalid response type from single flight")
	}

	return webfinger, nil
}

// FetchWellknownWebFinger retrieves wellknown web finger resource from a federated instance.
func (f *FediHelper) FetchWellknownWebFinger(ctx context.Context, serverHostname, username, domain string) (*WebFinger, error) {
	l := logger.WithField("func", "FetchWellknownWebFinger")

	hostMeta, err := f.FetchHostMeta(ctx, serverHostname)
	if err != nil {
		l.Errorf("nodeinfo: %s", err.Error())

		return nil, err
	}

	wfURI := hostMeta.WebfingerURI()
	if wfURI == "" {
		l.Errorf("host meta missing web finger url")

		return nil, NewError("host meta missing web finger url")
	}

	webfinger, err := f.FetchWebFinger(ctx, wfURI, username, domain)
	if err != nil {
		l.Errorf("webfinger: %s", err.Error())

		return nil, err
	}

	return webfinger, nil
}
