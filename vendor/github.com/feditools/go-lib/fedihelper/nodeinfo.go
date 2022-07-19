package fedihelper

import (
	"context"
	"encoding/json"
	"net/url"
)

// NodeInfoV2 is a federated node info 2.0 object.
type NodeInfoV2 struct {
	Metadata          map[string]interface{} `json:"metadata"`
	OpenRegistrations bool                   `json:"openRegistrations"`
	Protocols         []string               `json:"protocols"`
	Services          Services               `json:"services"`
	Software          Software               `json:"software"`
	Usage             Usage                  `json:"usage"`
	Version           string                 `json:"version"`
}

// Software contains the software and version of the node
type Software struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// Services contains the supported services of the node
type Services struct {
	Inbound  []string `json:"inbound"`
	Outbound []string `json:"outbound"`
}

// Usage contains usage statistics
type Usage struct {
	LocalPosts int64      `json:"localPosts"`
	Users      UsageUsers `json:"users"`
}

// UsageUsers contains usage statistics about users
type UsageUsers struct {
	Total int64 `json:"total"`
}

// findNodeInfo20URI parses a nodeinfo document for a nodeinfo 2.0 uri.
func findNodeInfo20URI(nodeinfo *WellKnownNodeInfo) (*url.URL, error) {
	var nodeinfoURIstr string
	for _, link := range nodeinfo.Links {
		if link.Rel == NodeInfo20Schema {
			nodeinfoURIstr = link.Href

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
func (f *FediHelper) GetNodeInfo20(ctx context.Context, domain string, infoURI *url.URL) (*NodeInfoV2, error) {
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

		// do request
		bodyBytes, err := f.http.InstanceGet(ctx, infoURI)
		if err != nil {
			l.Errorf("http get: %s", err.Error())

			return nil, err
		}

		// write cache
		err = f.kv.SetFediNodeInfo(ctx, domain, bodyBytes, f.nodeinfoCacheExp)
		if err != nil {
			fhErr := NewErrorf("redis set: %s", err.Error())
			l.Error(fhErr.Error())

			return nil, fhErr
		}

		return unmarshalNodeInfo20(bodyBytes)
	})

	if err != nil {
		fhErr := NewErrorf("singleflight: %s", err.Error())
		l.Error(fhErr.Error())

		return nil, fhErr
	}

	nodeinfo, ok := v.(*NodeInfoV2)
	if !ok {
		return nil, NewError("invalid response type from single flight")
	}

	return nodeinfo, nil
}

func unmarshalNodeInfo20(body []byte) (*NodeInfoV2, error) {
	var nodeinfo *NodeInfoV2
	if err := json.Unmarshal(body, &nodeinfo); err != nil {
		return nil, NewErrorf("unmarshal: %s", err.Error())
	}

	return nodeinfo, nil
}
