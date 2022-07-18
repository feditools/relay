package fedihelper

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/feditools/go-lib/fedihelper/models"
)

func (f *FediHelper) GetHostMeta(ctx context.Context, domain string) (*models.HostMeta, error) {
	l := logger.WithField("func", "GetHostMeta")

	hostmetaURI := fmt.Sprintf("https://%s/.well-known/host-meta", domain)
	v, err, _ := f.requestGroup.Do(hostmetaURI, func() (interface{}, error) {
		// do request
		resp, err := f.http.Get(ctx, hostmetaURI)
		if err != nil {
			l.Errorf("http get: %s", err.Error())

			return nil, err
		}

		hostMeta := new(models.HostMeta)
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			l.Errorf("read body: %s", err.Error())

			return nil, err
		}

		err = xml.Unmarshal(bodyBytes, hostMeta)
		if err != nil {
			l.Errorf("decode xml: %s", err.Error())

			return nil, err
		}

		return hostMeta, nil
	})

	if err != nil {
		l.Errorf("singleflight: %s", err.Error())

		return nil, err
	}

	hostMeta, ok := v.(*models.HostMeta)
	if !ok {
		return nil, NewError("invalid response type from single flight")
	}

	return hostMeta, nil
}

func (*FediHelper) WebfingerURIFromHostMeta(hostMeta *models.HostMeta) (string, error) {
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
