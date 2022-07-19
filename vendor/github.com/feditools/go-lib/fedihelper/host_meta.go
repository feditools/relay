package fedihelper

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/url"
	"strings"
)

type HostMeta struct {
	XMLNS string `xml:"xmlns,attr"`
	Links []Link `xml:"Link"`
}

func (h *HostMeta) WebfingerURI() WebfingerURI {
	for _, link := range h.Links {
		if link.Rel == HostMetaWebFingerTemplateRel {
			return WebfingerURI(link.Template)
		}
	}
	return ""
}

func (f *FediHelper) FetchHostMeta(ctx context.Context, domain string) (*HostMeta, error) {
	log := logger.WithField("func", "FetchHostMeta")
	hostMetaURI := &url.URL{
		Scheme: "https",
		Host:   domain,
		Path:   "/.well-known/host-meta",
	}

	v, err, _ := f.requestGroup.Do(fmt.Sprintf("hostmeta-%s", hostMetaURI.String()), func() (interface{}, error) {
		// do request
		bodyBytes, err := f.http.InstanceGet(ctx, hostMetaURI)
		if err != nil {
			log.Errorf("http get: %s", err.Error())

			return nil, err
		}

		hostMeta := new(HostMeta)
		err = xml.Unmarshal(bodyBytes, hostMeta)
		if err != nil {
			log.Errorf("decode xml: %s", err.Error())

			return nil, err
		}

		return hostMeta, nil
	})

	if err != nil {
		log.Errorf("singleflight: %s", err.Error())

		return nil, err
	}

	hostMeta, ok := v.(*HostMeta)
	if !ok {
		return nil, NewError("invalid response type from single flight")
	}

	return hostMeta, nil
}

func (*FediHelper) WebfingerURIFromHostMeta(hostMeta *HostMeta) (string, error) {
	// l := logger.WithField("func", "GetWebfingerURI")

	var hostMetaURITemplate string
	for _, link := range hostMeta.Links {
		if link.Rel == HostMetaWebFingerTemplateRel {
			hostMetaURITemplate = link.Template

			break
		}
	}
	if hostMetaURITemplate == "" {
		return "", NewError("web finger template not found")
	}

	// replace url template with golang f template
	if !strings.Contains(hostMetaURITemplate, "{uri}") {
		return "", NewError("web finger template invalid format")
	}
	hostMetaURIFTemplate := strings.ReplaceAll(hostMetaURITemplate, "{uri}", "acct:%s@%s")

	return hostMetaURIFTemplate, nil
}
