package fedihelper

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

// WellKnownNodeInfo is a federated well known node info object.
type WellKnownNodeInfo struct {
	Links []Link `json:"links"`
}

// GetWellknownNodeInfo retrieves wellknown nodeinfo from a federated instance.
func (f *FediHelper) GetWellknownNodeInfo(ctx context.Context, domain string) (*WellKnownNodeInfo, error) {
	l := logger.WithField("func", "GetWellknownNodeInfo")
	nodeInfoURI := &url.URL{
		Scheme: "https",
		Host:   domain,
		Path:   "/.well-known/nodeinfo",
	}

	v, err, _ := f.requestGroup.Do(fmt.Sprintf("hostmeta-%s", nodeInfoURI.String()), func() (interface{}, error) {
		// do request
		body, err := f.http.InstanceGet(ctx, nodeInfoURI)
		if err != nil {
			l.Errorf("http get: %s", err.Error())

			return nil, err
		}

		nodeinfo := new(WellKnownNodeInfo)
		err = json.Unmarshal(body, &nodeinfo)
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

	nodeinfo, ok := v.(*WellKnownNodeInfo)
	if !ok {
		return nil, NewError("invalid response type from single flight")
	}

	return nodeinfo, nil
}
