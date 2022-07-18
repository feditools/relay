package fedihelper

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/feditools/go-lib/fedihelper/models"
	libhttp "github.com/feditools/go-lib/http"
)

// GetWellknownNodeInfo retrieves wellknown nodeinfo from a federated instance.
func (f *FediHelper) GetWellknownNodeInfo(ctx context.Context, domain string) (*models.NodeInfo, error) {
	l := logger.WithField("func", "GetWellknownNodeInfo")
	nodinfoURI := fmt.Sprintf("https://%s/.well-known/nodeinfo", domain)
	v, err, _ := f.requestGroup.Do(nodinfoURI, func() (interface{}, error) {
		// do request
		resp, err := f.http.Get(ctx, nodinfoURI)
		if err != nil {
			l.Errorf("http get: %s", err.Error())

			return nil, err
		}

		nodeinfo := new(models.NodeInfo)
		defer resp.Body.Close()
		err = json.NewDecoder(resp.Body).Decode(nodeinfo)
		if err != nil {
			l.Errorf("decode json: %s", err.Error())

			return nil, err
		}

		return nodeinfo, nil
	})

	if err != nil {
		l.Errorf("singleflight: %s", err.Error())

		return nil, err
	}

	nodeinfo, ok := v.(*models.NodeInfo)
	if !ok {
		return nil, NewError("invalid response type from single flight")
	}

	return nodeinfo, nil
}

// GetWellknownWebFinger retrieves wellknown web finger resource from a federated instance.
func (f *FediHelper) GetWellknownWebFinger(ctx context.Context, serverHostname, username, domain string) (*models.WebFinger, error) {
	l := logger.WithField("func", "GetWellknownWebFinger")

	webfingerURI := fmt.Sprintf("https://%s/.well-known/webfinger?resource=acct:%%s@%%s", serverHostname)
	webfinger, err := f.webFinger(ctx, webfingerURI, username, domain)
	if err != nil {
		l.Errorf("singleflight: %s", err.Error())

		return nil, err
	}

	return webfinger, nil
}

// FindActorURI parses a webfinger document for an actor uri.
func FindActorURI(webfinger *models.WebFinger) (*url.URL, error) {
	var actorURIstr string
	for _, link := range webfinger.Links {
		if link.Rel == "self" || link.Type == string(libhttp.MimeAppActivityJSON) {
			actorURIstr = link.HRef

			break
		}
	}
	if actorURIstr == "" {
		return nil, nil
	}

	actorURI, err := url.Parse(actorURIstr)
	if err != nil {
		return nil, NewErrorf("invalid actor uri: %s", err.Error())
	}

	return actorURI, err
}
