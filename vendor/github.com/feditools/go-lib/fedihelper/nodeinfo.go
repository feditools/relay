package fedihelper

import (
	"context"
	"encoding/json"
	"io"
	nethttp "net/http"
	"net/url"

	"github.com/feditools/go-lib/fedihelper/models"
)

// findNodeInfo20URI parses a nodeinfo document for a nodeinfo 2.0 uri.
func findNodeInfo20URI(nodeinfo *models.NodeInfo) (*url.URL, error) {
	var nodeinfoURIstr string
	for _, link := range nodeinfo.Links {
		if link.Rel == NodeInfo20Schema {
			nodeinfoURIstr = link.HRef

			break
		}
	}
	if nodeinfoURIstr == "" {
		return nil, nil
	}

	nodeinfoURI, err := url.Parse(nodeinfoURIstr)
	if err != nil {
		return nil, NewErrorf("invalid nodeinfo 2.0 uri: %s", err.Error())
	}

	return nodeinfoURI, err
}

// GetNodeInfo20 retrieves wellknown nodeinfo from a federated instance.
func (f *FediHelper) GetNodeInfo20(ctx context.Context, domain string, infoURI *url.URL) (*models.NodeInfo2, error) {
	l := logger.WithField("func", "GetNodeInfo20")
	v, err, _ := f.requestGroup.Do(infoURI.String(), func() (interface{}, error) {
		// check cache
		cache, err := f.kv.GetFediNodeInfo(ctx, domain)
		if err != nil && err.Error() != "nil" {
			fhErr := NewErrorf("redis get: %s", err.Error())
			l.Error(fhErr.Error())

			return nil, fhErr
		}
		if err == nil {
			return unmarshalNodeInfo20(cache)
		}

		// get nodeinfo
		resp, err := f.http.Get(ctx, infoURI.String())
		if err != nil {
			fhErr := NewErrorf("http get: %s", err.Error())
			l.Error(fhErr.Error())

			return nil, fhErr
		}
		if resp.StatusCode != nethttp.StatusOK {
			return nil, NewErrorf("http status %s %d", infoURI, resp.StatusCode)
		}
		defer resp.Body.Close()
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fhErr := NewErrorf("read body: %s", err.Error())
			l.Error(fhErr.Error())

			return nil, fhErr
		}

		// write cache
		err = f.kv.SetFediNodeInfo(ctx, domain, bodyBytes, f.nodeinfoCacheExp)
		if err != nil {
			fhErr := NewErrorf("redis set: %s", err.Error())
			l.Error(fhErr.Error())

			return nil, fhErr
		}

		// marshal
		nodeinfo, err := unmarshalNodeInfo20(bodyBytes)
		if err != nil {
			fhErr := NewErrorf("marshal: %s", err.Error())
			l.Error(fhErr.Error())

			return nil, fhErr
		}

		return nodeinfo, nil
	})

	if err != nil {
		fhErr := NewErrorf("singleflight: %s", err.Error())
		l.Error(fhErr.Error())

		return nil, fhErr
	}

	nodeinfo, ok := v.(*models.NodeInfo2)
	if !ok {
		return nil, NewError("invalid response type from single flight")
	}

	return nodeinfo, nil
}

func unmarshalNodeInfo20(body []byte) (*models.NodeInfo2, error) {
	var nodeinfo *models.NodeInfo2
	if err := json.Unmarshal(body, &nodeinfo); err != nil {
		return nil, NewErrorf("unmarshal: %s", err.Error())
	}

	return nodeinfo, nil
}
